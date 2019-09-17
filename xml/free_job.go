package xml

import (
	"html/template"
	"bytes"
	"encoding/json"
)

var freeTestStyleJob = `
<project>
    <actions/>
    <description/>
    <keepDependencies>false</keepDependencies>
    <properties>
        <hudson.model.ParametersDefinitionProperty>
            <parameterDefinitions>
                <hudson.model.TextParameterDefinition>
                    <name>TAGNAME</name>
                    <description/>
                    <defaultValue/>
                </hudson.model.TextParameterDefinition>
            </parameterDefinitions>
        </hudson.model.ParametersDefinitionProperty>
    </properties>
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.8.0">
        <configVersion>2</configVersion>
        <userRemoteConfigs>
            <hudson.plugins.git.UserRemoteConfig>
                <url>{{.git_addr}}</url>
            </hudson.plugins.git.UserRemoteConfig>
        </userRemoteConfigs>
        <branches>
            <hudson.plugins.git.BranchSpec>
                <name>*/{{.git_type}}</name>
            </hudson.plugins.git.BranchSpec>
        </branches>
        <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
        <submoduleCfg class="list"/>
        <extensions/>
    </scm>
    <canRoam>true</canRoam>
    <disabled>false</disabled>
    <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
    <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
    <triggers/>
    <concurrentBuild>false</concurrentBuild>
    <builders>
        <hudson.tasks.Shell>
            <command>
                {{.command}}
            </command>
        </hudson.tasks.Shell>
    </builders>
    <publishers/>
    <buildWrappers/>
</project>
`

var freeStyleJob = `
<project>
    <actions/>
    <description/>
    <keepDependencies>false</keepDependencies>
    <properties>
        <hudson.security.AuthorizationMatrixProperty>
            <inheritanceStrategy class="org.jenkinsci.plugins.matrixauth.inheritance.InheritParentStrategy"/>
            <permission>hudson.model.Item.Build:{{.username}}</permission>
        </hudson.security.AuthorizationMatrixProperty>
        <com.sonyericsson.rebuild.RebuildSettings plugin="rebuild@1.28">
            <autoRebuild>false</autoRebuild>
            <rebuildDisabled>false</rebuildDisabled>
        </com.sonyericsson.rebuild.RebuildSettings>
        <hudson.model.ParametersDefinitionProperty>
            <parameterDefinitions>
                <net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition plugin="git-parameter@0.9.2">
                    <name>TAGNAME</name>
                    <description/>
                    <uuid>{{.user_token}}</uuid>
                    <type>PT_TAG</type>
                    <branch/>
                    <tagFilter>*</tagFilter>
                    <branchFilter>.*</branchFilter>
                    <sortMode>DESCENDING</sortMode>
                    <defaultValue/>
                    <selectedValue>NONE</selectedValue>
                    <quickFilterEnabled>false</quickFilterEnabled>
                    <listSize>5</listSize>
                </net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition>
            </parameterDefinitions>
        </hudson.model.ParametersDefinitionProperty>
    </properties>
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.8.0">
        <configVersion>2</configVersion>
        <userRemoteConfigs>
            <hudson.plugins.git.UserRemoteConfig>
                <url>{{.git_addr}}</url>
                <credentialsId>{{.git_token}}</credentialsId>
            </hudson.plugins.git.UserRemoteConfig>
        </userRemoteConfigs>
        <branches>
            <hudson.plugins.git.BranchSpec>
                <name>tags/$TAGNAME</name>
            </hudson.plugins.git.BranchSpec>
        </branches>
        <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
        <submoduleCfg class="list"/>
        <extensions/>
    </scm>
    <canRoam>true</canRoam>
    <disabled>false</disabled>
    <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
    <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
    <triggers/>
    <concurrentBuild>false</concurrentBuild>
    <builders>
        <hudson.tasks.Shell>
            <command>
                {{.command}}
            </command>
        </hudson.tasks.Shell>
    </builders>
    <publishers/>
    <buildWrappers/>
</project>
`

type Param struct {
	Username  string
	UserToken string
	GitAddr   string
	GitToken  string
	GitType   string
}

func MakeFreeStyleJob(name string, param Param, command string) (rs []byte, err error) {
	par := map[string]interface{}{
		"username":   param.Username,
		"user_token": param.UserToken,
		"git_addr":   param.GitAddr,
		"git_token":  param.GitToken,
		"command":    command,
	}
	tmpl, err := template.New(name).Parse(freeStyleJob)
	if err != nil {
		return
	}

	var w bytes.Buffer
	err = tmpl.Execute(&w, par)
	if err != nil {
		return
	}
	paramContentJson, err := json.Marshal(par)
	if err != nil {
		return
	}
	p := make([]byte, (len(freeStyleJob)*2)+(len(string(paramContentJson))*2))
	n, err := w.Read(p)
	rs = p[:n]
	return
}

func MakeTestStyleJob(name string, param Param, command string) (rs []byte, err error) {
	par := map[string]interface{}{
		"username":   param.Username,
		"user_token": param.UserToken,
		"git_addr":   param.GitAddr,
		"git_token":  param.GitToken,
		"command":    command,
		"git_type":   param.GitType,
	}
	tmpl, err := template.New(name).Parse(freeTestStyleJob)
	if err != nil {
		return
	}

	var w bytes.Buffer
	err = tmpl.Execute(&w, par)
	if err != nil {
		return
	}
	paramContentJson, err := json.Marshal(par)
	if err != nil {
		return
	}
	p := make([]byte, (len(freeTestStyleJob)*2)+(len(string(paramContentJson))*2))
	n, err := w.Read(p)
	rs = p[:n]
	return
}
