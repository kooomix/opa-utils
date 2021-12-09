package gitregostore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDefaultGitRegoStore(t *testing.T) {

	gs := InitDefaultGitRegoStore(-1)
	if gs.Rules == nil {
		t.Errorf("failed to decode")
	}
}

func TestInitGitRegoStoreFromRelease(t *testing.T) {
	// baseUrl := "https://api.github.com/repos"
	// owner := "armosec"
	// repository := "regolibrary"
	// path := "releases"
	// tag := "51267909"
	// branch := "master"
	// frequency := 15
	// gs := InitGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
	// if gs.Rules == nil {
	// 	t.Errorf("failed to decode")
	// }
	// fmt.Println(gs.URL)
}

// func TestInitGitRegoStoreFromRepo(t *testing.T) {
// 	baseUrl := "https://api.github.com/repos"
// 	owner := "armosec"
// 	repository := "regolibrary"
// 	path := "git/trees"
// 	tag := ""
// 	branch := "dev"
// 	frequency := 1
// 	gs := InitGitRegoStore(baseUrl, owner, repository, path, tag, branch, frequency)
// 	if gs.Controls == nil {
// 		t.Errorf("failed to decode controls")
// 	}
// 	if gs.Frameworks == nil {
// 		t.Errorf("failed to decode frameworks")
// 	}
// 	if gs.Rules == nil {
// 		t.Errorf("failed to decode rules")
// 	}
// 	fmt.Println(gs.URL)
// 	time.Sleep(time.Duration(frequency) * time.Minute)
// 	if len(gs.Frameworks) > 4 {
// 		t.Errorf("failed to decode controls")
// 	}
// }

func TestGetPoliciesMethods(t *testing.T) {
	gs := InitDefaultGitRegoStore(-1)

	index := 0

	// Rules
	policies, err := gs.GetOPAPolicies()
	if err != nil || policies == nil {
		t.Errorf("failed to get all policies %v", err)
	}
	policiesNames, err := gs.GetOPAPoliciesNamesList()
	if err != nil || len(policiesNames) == 0 {
		t.Errorf("failed to get policies names list %v", err)
		return
	}
	policy, err := gs.GetOPAPolicyByName(policiesNames[index])
	if err != nil || policy == nil {
		t.Errorf("failed to get policy by name: '%s', %v", policiesNames[index], err)
	}
	// Controls
	controls, err := gs.GetOPAControls()
	if err != nil || controls == nil {
		t.Errorf("failed to get all controls %v", err)
	}
	controlsNames, err := gs.GetOPAControlsNamesList()
	if err != nil || len(controlsNames) == 0 {
		t.Errorf("failed to get controls names list %v", err)
		return
	}

	control, err := gs.GetOPAControlByName(controlsNames[index])
	if err != nil || control == nil {
		t.Errorf("failed to get control by name: '%s', %v", controlsNames[index], err)
	}
	controlsIDs, err := gs.GetOPAControlsIDsList()
	if err != nil || len(controlsIDs) == 0 {
		t.Errorf("failed to get controls ids list %v", err)
		return
	}

	control, err = gs.GetOPAControlByID(controlsIDs[index])
	if err != nil || control == nil {
		t.Errorf("failed to get control by ID: '%s', %v", controlsNames[index], err)
	}
	// Frameworks
	frameworks, err := gs.GetOPAFrameworks()
	if err != nil || frameworks == nil {
		t.Errorf("failed to get all frameworks %v", err)
	}
	frameworksNames, err := gs.GetOPAFrameworksNamesList()
	if err != nil || len(frameworksNames) == 0 {
		t.Errorf("failed to get frameworks names list %v", err)
		return
	}
	framework, err := gs.GetOPAFrameworkByName(frameworksNames[0])
	if err != nil || framework == nil {
		t.Errorf("failed to get framework by name: '%s', %v", frameworksNames[0], err)
	}
}

func TestGetPoliciesMethodsNew(t *testing.T) {
	gs := NewDefaultGitRegoStore(-1)
	err := gs.SetRegoObjects()
	if err != nil {
		t.Errorf("error in SetRegoObjects: %v", err)
	}
	index := 0

	// Rules
	policies, err := gs.GetOPAPolicies()
	if err != nil || policies == nil {
		t.Errorf("failed to get all policies %v", err)
	}
	policiesNames, err := gs.GetOPAPoliciesNamesList()
	if err != nil || len(policiesNames) == 0 {
		t.Errorf("failed to get policies names list %v", err)
		return
	}
	policy, err := gs.GetOPAPolicyByName(policiesNames[index])
	if err != nil || policy == nil {
		t.Errorf("failed to get policy by name: '%s', %v", policiesNames[index], err)
	}
	// Controls
	controls, err := gs.GetOPAControls()
	if err != nil || controls == nil {
		t.Errorf("failed to get all controls %v", err)
	}
	controlsNames, err := gs.GetOPAControlsNamesList()
	if err != nil || len(controlsNames) == 0 {
		t.Errorf("failed to get controls names list %v", err)
		return
	}

	control, err := gs.GetOPAControlByName(controlsNames[index])
	if err != nil || control == nil {
		t.Errorf("failed to get control by name: '%s', %v", controlsNames[index], err)
	}
	controlsIDs, err := gs.GetOPAControlsIDsList()
	if err != nil || len(controlsIDs) == 0 {
		t.Errorf("failed to get controls ids list %v", err)
		return
	}

	control, err = gs.GetOPAControlByID(controlsIDs[index])
	if err != nil || control == nil {
		t.Errorf("failed to get control by ID: '%s', %v", controlsNames[index], err)
	}
	// Frameworks
	frameworks, err := gs.GetOPAFrameworks()
	if err != nil || frameworks == nil {
		t.Errorf("failed to get all frameworks %v", err)
	}
	frameworksNames, err := gs.GetOPAFrameworksNamesList()
	if err != nil || len(frameworksNames) == 0 {
		t.Errorf("failed to get frameworks names list %v", err)
		return
	}
	framework, err := gs.GetOPAFrameworkByName(frameworksNames[0])
	if err != nil || framework == nil {
		t.Errorf("failed to get framework by name: '%s', %v", frameworksNames[0], err)
	}
}

func TestIsControlID(t *testing.T) {
	assert.True(t, isControlID("c-0001"))
	assert.True(t, isControlID("C-0001"))
	assert.False(t, isControlID("c-00012"))
	assert.False(t, isControlID("t-0001"))
	assert.False(t, isControlID("my name"))
}
