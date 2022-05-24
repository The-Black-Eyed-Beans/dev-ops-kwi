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
	Vpc     VpcJson    `json:"vpc"`
	Public  SubnetJson `json:"public"`
	Private SubnetJson `json:"private"`
	AppLB   AppJson    `json:"alb"`
	GateLB  GateJson   `json:"glb"`
	Ecs     EcsJson    `json:"ecs"`
	Eks     EksJson    `json:"eks"`
}

type VpcJson struct {
	Name string `json:"name"`
	Cidr string `json:"cidr"`
}

type SubnetJson struct {
	Name string `json:"name"`
	Cidr string `json:"cidr"`
}

type AppJson struct {
	Name string `json:"name"`
}

type GateJson struct {
	Name string `json:"name"`
}

type EcsJson struct {
	Name string `json:"name"`
}

type EksJson struct {
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
			Name: terraform.Output(t, terraformOptions, "vpc_name"),
			Cidr: terraform.Output(t, terraformOptions, "vpc_cidr"),
		},
		Public: SubnetJson{
			Name: terraform.Output(t, terraformOptions, "public_one_name"),
			Cidr: terraform.Output(t, terraformOptions, "public_one_cidr"),
		},
		Private: SubnetJson{
			Name: terraform.Output(t, terraformOptions, "private_two_name"),
			Cidr: terraform.Output(t, terraformOptions, "private_two_cidr"),
		},
		AppLB: AppJson{
			Name: terraform.Output(t, terraformOptions, "app_load_balancer_name"),
		},
		GateLB: GateJson{
			Name: terraform.Output(t, terraformOptions, "gateway_load_balancer_name"),
		},
		Ecs: EcsJson{
			Name: terraform.Output(t, terraformOptions, "ecs_name"),
		},
		Eks: EksJson{
			Name: terraform.Output(t, terraformOptions, "eks_name"),
		},
	}

	if assert.Equal(t, ExpectedJson.AppLB, ActualJson.AppLB) {
		deployment_passed = true
		t.Logf("PASS: The expected App Load Balancer name, %v, matches the actual App Load Balancer name.", ExpectedJson.AppLB.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v", ExpectedJson.AppLB.Name, ActualJson.AppLB.Name)
	}

	if assert.Equal(t, ExpectedJson.GateLB, ActualJson.GateLB) {
		deployment_passed = true
		t.Logf("PASS: The expected Gateway Load Balancer name, %v, matches the actual Gateway Load Balancer name.", ExpectedJson.GateLB.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.GateLB.Name, ActualJson.GateLB.Name)
	}

	if assert.Equal(t, ExpectedJson.Eks, ActualJson.Eks) {
		deployment_passed = true
		t.Logf("PASS: The expected EKS cluster name, %v, matches the actual EKS cluster name.", ExpectedJson.Eks.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.Eks.Name, ActualJson.Eks.Name)
	}

	if assert.Equal(t, ExpectedJson.Ecs, ActualJson.Ecs) {
		deployment_passed = true
		t.Logf("PASS: The expected ECS cluster name, %v, matches the actual ECS cluster name.", ExpectedJson.Ecs.Name)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.Ecs.Name, ActualJson.Ecs.Name)
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

	if assert.Equal(t, ExpectedJson.Public, ActualJson.Public) {
		deployment_passed = true
		t.Logf("PASS: The expected name of the first public subnet, %v, matches the actual name of the first public subnet.", ExpectedJson.Public.Name)
		t.Logf("PASS: The expected CIDR block for the first public subnet, %v, matches the actual CIDR block for the first public subnet.", ExpectedJson.Public.Cidr)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.Public, ActualJson.Public)
	}

	if assert.Equal(t, ExpectedJson.Private, ActualJson.Private) {
		deployment_passed = true
		t.Logf("PASS: The expected name of the second private subnet, %v, matches the actual name of the second private subnet.", ExpectedJson.Private.Name)
		t.Logf("PASS: The expected CIDR block for the second private subnet, %v, matches the actual CIDR block for the second private subnet.", ExpectedJson.Private.Cidr)
	} else {
		deployment_passed = false
		terraform.Destroy(t, terraformOptions)
		t.Fatalf("FAIL: Expected %v, but found %v.", ExpectedJson.Private, ActualJson.Private)
	}

	time.Sleep(120 * time.Second)
	fmt.Println("Sleep Over...")
}
