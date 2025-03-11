//go:build !integration

package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mainTestSuite struct {
	suite.Suite
	// region variables

	//endregion variables
}

// region setup
func (suite *mainTestSuite) SetupSuite()                           {}
func (suite *mainTestSuite) TearDownSuite()                        {}
func (suite *mainTestSuite) BeforeTest(suiteName, testName string) {}
func (suite *mainTestSuite) AfterTest(suiteName, testName string)  {}
func (suite *mainTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	if !stats.Passed() {
		buf := strings.Builder{}
		for _, information := range stats.TestStats {
			if !information.Passed {
				buf.WriteString(fmt.Sprintf("Failed %s.%s\n", suiteName, information.TestName))
			}
		}
		suite.Fail(buf.String())
	}
}
func TestRunSuitemain(t *testing.T) {
	suite.Run(t, new(mainTestSuite))
}

// endregion setup
// region tests
func (suite *mainTestSuite) TestFileNameIsCorrect() {
	dir, err := getDirectory([]string{"testdata"})
	suite.Require().NoError(err)
	suite.Require().NotNil(dir)
	stat, err := dir.Stat()
	suite.Require().NoError(err)
	suite.Require().NotNil(stat)
	suite.Require().True(stat.IsDir())
}

func (suite *mainTestSuite) TestGetImage() {
	img, err := GetImage([]string{"testdata/allfiles"})
	suite.Require().NoError(err)
	suite.Require().NotNil(img)
	file, err := os.Open(img)
	defer file.Close()

	suite.Require().Contains([]string{"0-zero", "1-one", "2-two"}, filepath.Base(file.Name()))
}

func (suite *mainTestSuite) TestUsingDot() {
	img, err := GetImage([]string{"."})
	suite.Require().NoError(err)
	suite.Require().NotNil(img)
	file, err := os.Open(img)
	defer file.Close()
	suite.Require().Contains([]string{"getBackgroundImage.go", "getBackgroundImage_test.go"}, filepath.Base(file.Name()))
}

func (suite *mainTestSuite) TestGetDirWithOnlyDirs() {
	file, err := GetImage([]string{"testdata/onlydirs"})
	suite.Require().Equal("", file)
	suite.Require().Error(err)
	suite.Require().Equal("No files in directory", err.Error())
}

// endregion tests
