package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/tokinaa/warpcast-tools/warpcast"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConfigStruct struct {
	Accounts          []string `json:"accounts"`
	DelayFollow       int      `json:"delayFollow"`
	DelayUnfollow     int      `json:"delayUnfollow"`
	DelayLike         int      `json:"delayLike"`
	DelayComment      int      `json:"delayComment"`
	DelayRecast       int      `json:"delayRecast"`
	DelayTimeline     int      `json:"delayTimeline"`
	CustomCommentText []string `json:"customCommentText"`
}

var (
	myConfig = LoadConfig()
)

func LoadConfig() ConfigStruct {
	// Load from config.json
	openFile, err := os.Open("config.json")
	if err != nil {
		return ConfigStruct{}
	}

	defer openFile.Close()

	var config ConfigStruct
	jsonParser := json.NewDecoder(openFile)
	jsonParser.Decode(&config)

	return config
}

func init() {
	openLoadConfig := LoadConfig()
	myConfig = openLoadConfig
}

func checkingError(err error) {
	if err != nil {
		switch {
		case err.Error() == "interrupt":
			os.Exit(0)
		default:
			break
		}
	}
}

func showPressEnter() {
	fmt.Print("Press Enter to Back...")

	var input string
	fmt.Scanln(&input)
}

func multiAccountsManagement() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Multi Accounts Management")
	fmt.Println("Total Accounts :", len(myConfig.Accounts))
	fmt.Println()
	fmt.Println("1. List Account")
	fmt.Println("2. Add Account")
	fmt.Println("3. Remove Accounts")
	fmt.Println("4. Back")
	fmt.Println()

	inputMenu := ""
	inputMenuError := survey.AskOne(&survey.Input{
		Message: "Select Menu:",
	}, &inputMenu, survey.WithValidator(survey.Required))

	checkingError(inputMenuError)

	switch inputMenu {
	case "1":
		fmt.Print("\033[H\033[2J")
		fmt.Println("List Account")
		fmt.Println()
		for i, account := range myConfig.Accounts {
			fmt.Println(i+1, account)
		}
		fmt.Println()

		showPressEnter()

		fmt.Print("\033[H\033[2J")
		multiAccountsManagement()
		break
	case "2":
		fmt.Print("\033[H\033[2J")
		fmt.Println("Add Account")
		fmt.Println()

		inputAccount := ""
		inputAccountError := survey.AskOne(&survey.Input{
			Message: "Authorization Token:",
		}, &inputAccount, survey.WithValidator(survey.Required))

		checkingError(inputAccountError)

		myConfig.Accounts = append(myConfig.Accounts, inputAccount)

		file, _ := json.MarshalIndent(myConfig, "", " ")
		_ = os.WriteFile("config.json", file, 0644)

		fmt.Println("Account Added")
		fmt.Println()

		showPressEnter()

		fmt.Print("\033[H\033[2J")
		multiAccountsManagement()
		break
	case "3":
		fmt.Print("\033[H\033[2J")
		fmt.Println("Remove Accounts")
		fmt.Println()

		for i, account := range myConfig.Accounts {
			fmt.Println(i+1, account)
		}

		fmt.Println()

		inputAccount := 0
		inputAccountError := survey.AskOne(&survey.Select{
			Message: "Select Account:",
			Options: myConfig.Accounts,
		}, &inputAccount, survey.WithValidator(survey.Required))

		checkingError(inputAccountError)

		myConfig.Accounts = append(myConfig.Accounts[:inputAccount], myConfig.Accounts[inputAccount+1:]...)

		file, _ := json.MarshalIndent(myConfig, "", " ")
		_ = os.WriteFile("config.json", file, 0644)

		fmt.Println("Account Removed")
		fmt.Println()

		showPressEnter()

		fmt.Print("\033[H\033[2J")
		multiAccountsManagement()
		break
	case "4":
		fmt.Print("\033[H\033[2J")
		main()
		break
	}
}

func autoTimeline() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Auto Like, Comment, and Recast Timeline")
	fmt.Println()

	inputSelectAccount := 0
	inputSelectAccountError := survey.AskOne(&survey.Select{
		Message: "Select Account:",
		Options: myConfig.Accounts,
	}, &inputSelectAccount, survey.WithValidator(survey.Required))

	checkingError(inputSelectAccountError)

	inputSelectMode := []string{}
	inputSelectModeError := survey.AskOne(&survey.MultiSelect{
		Message: "Select Mode:",
		Options: []string{"Like", "Comments", "Recast"},
	}, &inputSelectMode, survey.WithValidator(survey.Required))

	checkingError(inputSelectModeError)

	inputChoiceTimeline := ""
	inputChoiceTimelineError := survey.AskOne(&survey.Select{
		Message: "Select Timeline:",
		Options: []string{"Home", "All-Channels"},
	}, &inputChoiceTimeline, survey.WithValidator(survey.Required))

	checkingError(inputChoiceTimelineError)

	fmt.Println()

	var excludeHash []string
	var lastTimestamp int64 = 0

	for {
		timeline, err := warpcast.GetFeedsItems(myConfig.Accounts[inputSelectAccount], inputChoiceTimeline, lastTimestamp, excludeHash)
		if err != nil {
			fmt.Printf("[TIMELINE][GETTER] ERROR : %s\n", err)
			break
		}

		if lastTimestamp == 0 {
			lastTimestamp = timeline.Result.LatestMainCastTimestamp
		}

		items := timeline.Result.Items

		if len(items) == 0 {
			delayTimeline := time.Duration(myConfig.DelayTimeline) * time.Millisecond
			time.Sleep(delayTimeline)
			continue
		}

		for _, item := range items {
			if !strings.Contains(strings.Join(excludeHash, ","), item.Cast.Hash[2:10]) {
				excludeHash = append(excludeHash, item.Cast.Hash[2:10])
			}

			fmt.Printf("[TIMELINE] [%s] ", item.Cast.Hash)

			// Check if Like in inputSelectMode
			if strings.Contains(strings.Join(inputSelectMode, ","), "Like") {
				fmt.Printf("[LIKE]")

				if item.Cast.ViewerContext.Reacted {
					fmt.Printf(" ALREADY ")
				} else {
					_, err := warpcast.Like(myConfig.Accounts[inputSelectAccount], item.Cast.Hash)
					if err != nil {
						fmt.Printf(" ERROR : %s", err)
					} else {
						fmt.Printf(" SUCCESS")
					}
					fmt.Printf(" ")

					delayLike := time.Duration(myConfig.DelayLike) * time.Millisecond
					time.Sleep(delayLike)
				}
			}

			// Check if Comment in inputSelectMode
			if strings.Contains(strings.Join(inputSelectMode, ","), "Comments") {
				fmt.Printf("[COMMENT]")

				commentText := ""
				if strings.Contains(item.Cast.Text, "$DEGEN") {
					randomThreeDigit := rand.Intn(999-100+1) + 100
					commentText = fmt.Sprintf("%d $DEGEN", randomThreeDigit)
				}

				if commentText != "" {
					_, err := warpcast.Comment(myConfig.Accounts[inputSelectAccount], item.Cast.Hash, commentText)
					if err != nil {
						fmt.Printf(" ERROR : %s", err)
					} else {
						fmt.Printf(" SUCCESS [%s]", commentText)
					}
				}

				fmt.Printf(" ")
			}

			// Check if Recast in inputSelectMode
			if strings.Contains(strings.Join(inputSelectMode, ","), "Recast") {
				fmt.Printf("[RECAST]")

				if item.Cast.ViewerContext.Recast {
					fmt.Printf(" ALREADY ")
				} else {
					_, err := warpcast.Recast(myConfig.Accounts[inputSelectAccount], item.Cast.Hash)
					if err != nil {
						fmt.Printf(" ERROR : %s", err)
					} else {
						fmt.Printf(" SUCCESS")
					}

					fmt.Printf(" ")

					delayRecast := time.Duration(myConfig.DelayRecast) * time.Millisecond
					time.Sleep(delayRecast)
				}
			}

			fmt.Printf("\n")
		}

		fmt.Println()
		fmt.Printf("\tWaiting for %ds to get new feeds...\n", myConfig.DelayTimeline/1000)
		fmt.Println()

		delayTimeline := time.Duration(myConfig.DelayTimeline) * time.Millisecond
		time.Sleep(delayTimeline)
	}
}

func followTarget() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Follow Following/Followers Target")
	fmt.Println()

	inputSelectAccount := 0
	inputSelectAccountError := survey.AskOne(&survey.Select{
		Message: "Select Account:",
		Options: myConfig.Accounts,
	}, &inputSelectAccount, survey.WithValidator(survey.Required))

	checkingError(inputSelectAccountError)

	inputTargetUsername := ""
	inputTargetUsernameError := survey.AskOne(&survey.Input{
		Message: "Target Username:",
	}, &inputTargetUsername, survey.WithValidator(survey.Required))

	checkingError(inputTargetUsernameError)

	inputChoiceMode := ""
	inputChoiceModeError := survey.AskOne(&survey.Select{
		Message: "Select Mode:",
		Options: []string{"Following", "Followers"},
	}, &inputChoiceMode, survey.WithValidator(survey.Required))

	checkingError(inputChoiceModeError)

	fmt.Println()

	fmt.Printf("[%s] Getting Data of @%s...\n", inputChoiceMode, inputTargetUsername)

	profile, err := warpcast.GetProfile(myConfig.Accounts[inputSelectAccount], inputTargetUsername)
	if err != nil {
		fmt.Printf("[PROFILE][GETTER] ERROR : %s\n", err)
		return
	}

	fmt.Printf("[%s] [@%s] FID : %d | Followers : %d | Following : %d\n", inputChoiceMode, inputTargetUsername, profile.Result.User.Fid, profile.Result.User.FollowerCount, profile.Result.User.FollowingCount)
	fmt.Println()

	var cursor string = ""
	for {
		fidStr := strconv.Itoa(profile.Result.User.Fid)
		tryToGetFollowersOrFollowing, err := warpcast.GetProfileInformation(strings.ToLower(inputChoiceMode), myConfig.Accounts[inputSelectAccount], fidStr, cursor)
		if err != nil {
			fmt.Printf("[GET DATA][%s] FAILED TO GET DATA | ERROR : %s\n", inputChoiceMode, err)
			continue
		}
		for _, item := range tryToGetFollowersOrFollowing.Result.Users {
			fidTarget := strconv.Itoa(item.Fid)
			fmt.Printf("[%s] [@%s] FID : %s", inputChoiceMode, item.Username, fidTarget)

			if item.ViewerContext.Following {
				fmt.Printf(" SKIP YOU ALREADY FOLLOW !\n")
				continue
			}

			_, err := warpcast.Follow(myConfig.Accounts[inputSelectAccount], fidTarget)
			if err != nil {
				fmt.Printf(" ERROR : %s\n", err)
			} else {
				fmt.Printf(" SUCCESS\n")
			}

			delayFollow := time.Duration(myConfig.DelayFollow) * time.Millisecond
			time.Sleep(delayFollow)
		}

		if tryToGetFollowersOrFollowing.Next.Cursor == "" {
			break
		}

		cursor = tryToGetFollowersOrFollowing.Next.Cursor
	}
}

func unfollowNotFB() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Unfollow Who Not Follow Back")
	fmt.Println()

	inputSelectAccount := 0
	inputSelectAccountError := survey.AskOne(&survey.Select{
		Message: "Select Account:",
		Options: myConfig.Accounts,
	}, &inputSelectAccount, survey.WithValidator(survey.Required))

	checkingError(inputSelectAccountError)

	fmt.Println()
	fmt.Printf("[PROFILE] Getting following data\n")

	profile, err := warpcast.GetMyProfile(myConfig.Accounts[inputSelectAccount])
	if err != nil {
		fmt.Printf("[PROFILE][GETTER] ERROR : %s\n", err)
		return
	}

	fidStr := strconv.Itoa(profile.Result.State.User.Fid)

	fmt.Printf("[%s] [@%s] FID : %s | Followers : %d | Following : %d\n", "PROFILE", profile.Result.State.User.Username, fidStr, profile.Result.State.User.FollowerCount, profile.Result.State.User.FollowingCount)
	fmt.Println()

	var cursor string = ""
	for {
		tryToGetFollowing, err := warpcast.GetProfileInformation("following", myConfig.Accounts[inputSelectAccount], fidStr, cursor)
		if err != nil {
			fmt.Printf("[GET DATA][FOLLOWING] FAILED TO GET DATA | ERROR : %s\n", err)
			continue
		}
		for _, item := range tryToGetFollowing.Result.Users {
			fidTarget := strconv.Itoa(item.Fid)
			fmt.Printf("[UNFOLLOW] [@%s] FID : %s", item.Username, fidTarget)

			if item.ViewerContext.FollowedBy {
				fmt.Printf(" SKIP THEY FOLLOW YOU !\n")
				continue
			}

			_, err := warpcast.Unfollow(myConfig.Accounts[inputSelectAccount], fidTarget)
			if err != nil {
				fmt.Printf(" ERROR : %s\n", err)
			} else {
				fmt.Printf(" SUCCESS\n")
			}

			delayUnfollow := time.Duration(myConfig.DelayUnfollow) * time.Millisecond
			time.Sleep(delayUnfollow)
		}

		if tryToGetFollowing.Next.Cursor == "" {
			break
		}

		cursor = tryToGetFollowing.Next.Cursor
	}
}

func main() {
	fmt.Println("Warpcast Tools")
	fmt.Println("Author : @x0xdead / Wildaann")
	fmt.Println()
	fmt.Println("1. Multi Accounts Management")
	fmt.Println("2. Follow Target (Followers/Following)")
	fmt.Println("3. Auto Like, Comment, and Recast Timeline (Home/All-Channels)")
	fmt.Println("4. Unfollow Who Not Follow Back")
	fmt.Println()

	inputMenu := ""
	inputMenuError := survey.AskOne(&survey.Input{
		Message: "Select Menu:",
	}, &inputMenu, survey.WithValidator(survey.Required))

	checkingError(inputMenuError)

	switch inputMenu {
	case "1":
		multiAccountsManagement()
		break
	case "2":
		followTarget()
		break
	case "3":
		autoTimeline()
		break
	case "4":
		unfollowNotFB()
		break
	}
}
