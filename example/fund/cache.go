package main

import (
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"io/ioutil"
	"os"
)

type FundBasicInfo struct {
	Code      string
	ShortName string
	Name      string
	Type      string
	FullName  string
}

type FundCache struct {
	Length int
	List   []FundBasicInfo
}

func LoadCache(file string) (*FundCache, error) {

	log.Info("load cache ...")

	var cache FundCache

	fp, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	defer fp.Close()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(bytes) > 0 {
		err = util.Json.Unmarshal(bytes, &cache)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}

	return &cache, nil
}

func SaveCache(cache *FundCache, file string) error {
	log.Info("save cache ...")

	s, err := util.Json.MarshalIndent(cache, "", "  ")
	if err != nil {
		log.Error(err)
	}

	err = ioutil.WriteFile(file, s, 755)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
