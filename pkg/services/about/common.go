package about

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const aboutFile = "./data/about/about.json"

func init() {
	f, err := os.OpenFile(aboutFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if os.IsExist(err) {
		return
	}
	if err != nil {
		log.Fatalf("in about.init(): unexpected error while opening about.json: %v", err)
	}

	ai := newDefaultAboutInfo()

	en := json.NewEncoder(f)
	if err = en.Encode(ai); err != nil {
		log.Fatalf("in about.init(): unexpected error while trying to encode into about.json: %v", err)
	}

	defer f.Close()
}

func GetInfo() (AboutInfo, error) {
	var info AboutInfo
	f, err := os.OpenFile(aboutFile, os.O_RDWR, 0644)
	if err != nil {
		return info, err
	}
	defer f.Close()

	d := json.NewDecoder(f)

	if err = d.Decode(&info); err != nil {
		return info, err
	}

	return info, nil
}

func UpdateName(name string) error {
	var info AboutInfo

	// step 1: read current about info
	rf, err := os.Open(aboutFile)
	if err != nil {
		return fmt.Errorf("open for read: %w", err)
	}
	if err := json.NewDecoder(rf).Decode(&info); err != nil {
		rf.Close()
		return fmt.Errorf("decode failed: %w", err)
	}
	rf.Close()

	// step 2: update about info
	info.Name = name

	// step 3: write it back to file
	wf, err := os.OpenFile(aboutFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer wf.Close()

	if err := json.NewEncoder(wf).Encode(info); err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	return nil
}

func UpdateAvatarURL(url string) error {
	var info AboutInfo

	// step 1: read current about info
	rf, err := os.Open(aboutFile)
	if err != nil {
		return fmt.Errorf("open for read: %w", err)
	}
	if err := json.NewDecoder(rf).Decode(&info); err != nil {
		rf.Close()
		return fmt.Errorf("decode failed: %w", err)
	}
	rf.Close()

	// step 2: update about info
	info.AvatarURL = url

	// step 3: write it back to file
	wf, err := os.OpenFile(aboutFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer wf.Close()

	if err := json.NewEncoder(wf).Encode(info); err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	return nil
}

func UpdateShortDescription(desc string) error {
	var info AboutInfo

	// step 1: read current about info
	rf, err := os.Open(aboutFile)
	if err != nil {
		return fmt.Errorf("open for read: %w", err)
	}
	if err := json.NewDecoder(rf).Decode(&info); err != nil {
		rf.Close()
		return fmt.Errorf("decode failed: %w", err)
	}
	rf.Close()

	// step 2: update about info
	info.DescShort = desc

	// step 3: write it back to file
	wf, err := os.OpenFile(aboutFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer wf.Close()

	if err := json.NewEncoder(wf).Encode(info); err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	return nil
}

func UpdateDescription(desc string) error {
	var info AboutInfo

	// step 1: read current about info
	rf, err := os.Open(aboutFile)
	if err != nil {
		return fmt.Errorf("open for read: %w", err)
	}
	if err := json.NewDecoder(rf).Decode(&info); err != nil {
		rf.Close()
		return fmt.Errorf("decode failed: %w", err)
	}
	rf.Close()

	// step 2: update about info
	info.Description = desc

	// step 3: write it back to file
	wf, err := os.OpenFile(aboutFile, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("open for write: %w", err)
	}
	defer wf.Close()

	if err := json.NewEncoder(wf).Encode(info); err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	return nil
}
