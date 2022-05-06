package aws_infrastructure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

type MainJson struct {
	Vpc    VpcJson  `json:"vpc"`
	AppLB  AppJson  `json:"alb"`
	GateLB GateJson `json:"glb"`
}

type VpcJson struct {
	Name string `json:"name"`
	Cidr string `json:"cidr"`
}

type AppJson struct {
	Name string `json:"name"`
}

type GateJson struct {
	Name string `json:"name"`
}

var deployment_passed bool
var ExpectedJson MainJson

func init() {
	jsonFile, err := os.Open("./aws_infrastructure_testdata.json")
	if err != nil {
		fmt.Println(err)
	}
	byteVal, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteVal, &ExpectedJson)
}
func TestAWSInfraInput(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../deployment",
		VarFiles:     []string{"input-dev.tfvars"},
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	ActualJson := MainJson{
		Vpc: VpcJson{
			Name: terraform.Output(t, terraformOptions, "vpc"),
			Cidr: terraform.Output(t, terraformOptions, "vpc_cidr"),
		},
		AppLB: AppJson{
			Name: terraform.Output(t, terraformOptions, "app_load_balancer"),
		},
		GateLB: GateJson{
			Name: terraform.Output(t, terraformOptions, "gateway_load_balancer"),
		},
	}

	if assert.Equal(t, ExpectedJson.AppLB, ActualJson.AppLB) {
		deployment_passed = true
		t.Logf("PASS: The expected App Load Balancer name, %v, matches the actual App Load Balancer Name.", ExpectedJson.AppLB.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v", ExpectedJson.AppLB.Name, ActualJson.AppLB.Name)
	}

	if assert.Equal(t, ExpectedJson.GateLB, ActualJson.GateLB) {
		deployment_passed = true
		t.Logf("PASS: The expected Gateway Load Balancer name, %v, matches the actual Gateway Load Balancer Name.", ExpectedJson.GateLB.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.GateLB.Name, ActualJson.GateLB.Name)
	}

	if assert.Equal(t, ExpectedJson.Vpc, ActualJson.Vpc) {
		deployment_passed = true
		t.Logf("PASS: The expected name of the VPC, %v, matches the actual name of the VPC.", ExpectedJson.Vpc.Name)
		t.Logf("PASS: The expected CIDR block for the VPC, %v, matches the actual CIDR block for the VPC.", ExpectedJson.Vpc.Cidr)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.Vpc, ActualJson.Vpc)
	}

	time.Sleep(120 * time.Second)
	fmt.Println("Sleep Over...")
}
