/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 10:53:23
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-28 10:56:25
 * @FilePath: /dao-generator/cmd/errors.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cmd

import (
	"os"

	"go.uber.org/zap"
)

func CheckErrWithMsg(err error, msg string) {
	if err != nil {
		// if errors.Is(err, workspace.ErrGetHomeDir) ||
		// 	errors.Is(err, workspace.ErrCreateRepoDir) ||
		// 	errors.Is(err, workspace.ErrRepoPathIsFile) ||
		// 	errors.Is(err, workspace.ErrGlobalConfigDirNotFind) ||
		// 	errors.Is(err, workspace.ErrNotAbsPath) ||
		// 	errors.Is(err, workspace.ErrRepoConfigDirNotFind) ||
		// 	errors.Is(err, workspace.ErrOutsideRepoFile) ||
		// 	errors.Is(err, dvc.ErrNoRemoteURL) ||
		// 	errors.Is(err, dvc.ErrURLIllegal) ||
		// 	errors.Is(err, vcs.ErrBranchExisted) ||
		// 	errors.Is(err, dvc.ErrAllFilesDeleted) ||
		// 	errors.Is(err, dvc.ErrInvalidBranchName) ||
		// 	errors.Is(err, syscall.EACCES) ||
		// 	errors.Is(err, service.ErrRemoteRepoNotConfigured) ||
		// 	errors.Is(err, workspace.ErrCheckAnnotationFile) ||
		// 	errors.Is(err, workspace.ErrInvalidAnnotationFile) {
		// 	log.ConsoleW(err.Error())
		// } else if errors.Is(err, dvc.ErrChangeNotCommitted) {
		// 	log.ConsoleW("error: Your local changes would be overwritten by checkout")
		// 	log.ConsoleW("hint: use `grav status` to view the list of conflicting files")
		// } else if errors.Is(err, dvc.ErrNothingToCommit) {
		// 	log.ConsoleW("nothing to commit, working tree clean")
		// } else if errors.Is(err, dvc.ErrURLWithOldVersion) {
		// 	log.ConsoleW("This URL is an old version and no longer supported. \nPlease go to the web and copy the latest version of the URL.")
		// } else if errors.Is(err, dvc.ErrRemoteResourceNotFound) || errors.Is(err, service.ErrResourceNotFound) {
		// 	os.Exit(0)
		// } else if errors.Is(err, storage.ErrAossForbiddenAccess) {
		// 	log.ConsoleW("upload/download objects to aoss failed. Please check aoss credential (ak/sk) and aoss resource capacity.")
		// } else if errors.Is(err, dvc.ErrRemoteRepoHasGravDir) {
		// 	log.ConsoleW("The remote dataset has a '.grav' folder in the root path, clone canceled.")
		// } else if errors.Is(err, toolkit.ErrParseReadme) || errors.Is(err, toolkit.ErrTagSupport) || errors.Is(err, toolkit.ErrYamlDataSupport) {
		// 	log.ConsoleW("README.md YAML metadata error.")
		// } else if errors.Is(err, service.ErrAPIConnectError) {
		// 	log.ConsoleW("unable to connect the datahub service")
		// 	log.Infof("%s %s", "please try again or check error in", log.ErrLogFilePath)
		// } else if errors.Is(err, service.ErrForbidden) {
		// 	log.ConsoleW(workspace.Red + "No permission to access this dataset" + workspace.Reset)
		// } else if errors.Is(err, service.ErrRepoDataExceedLimit) {
		// 	log.ConsoleW(service.ErrRepoDataExceedLimit.Error())
		// } else if errors.Is(err, service.ErrPushTaskDuplicated) {
		// 	log.ConsoleW(service.ErrPushTaskDuplicated.Error())
		// } else if errors.Is(err, service.ErrRemoteRepoNotFound) {
		// 	log.ConsoleW(service.ErrRemoteRepoNotFound.Error())
		// } else if errors.Is(err, service.ErrResourceStateInvalid) {
		// 	log.ConsoleW(service.ErrResourceStateInvalid.Error())
		// } else if errors.Is(err, service.ErrResourceDowngrade) {
		// 	log.ConsoleW(service.ErrResourceDowngrade.Error())
		// } else if errors.Is(err, service.ErrDeadlineExceededError) {
		// 	log.ConsoleW(service.ErrDeadlineExceededError.Error())
		// } else if errors.Is(err, service.ErrAOssRequestParamsError) {
		// 	log.ConsoleW(service.ErrAOssRequestParamsError.Error())
		// } else if errors.Is(err, service.ErrUnauthenticated) {
		// 	log.ConsoleW(service.ErrUnauthenticated.Error())
		// } else if errors.Is(err, service.ErrAKSKNotConfigured) {
		// 	log.ConsoleW(service.ErrAKSKNotConfigured.Error())
		// } else if errors.Is(err, service.ErrTaskSourcePathNotFound) {
		// 	log.ConsoleW(service.ErrTaskSourcePathNotFound.Error())
		// } else if errors.Is(err, service.ErrUserTaskRepoIsEmpty) {
		// 	log.ConsoleW(service.ErrUserTaskRepoIsEmpty.Error())
		// } else if errors.Is(err, service.ErrTaskAllFileUploadFailed) {
		// 	log.ConsoleW(service.ErrTaskAllFileUploadFailed.Error())
		// } else if errors.Is(err, service.ErrUserTaskForbiddenWhenRepoScanning) {
		// 	log.ConsoleW(service.ErrUserTaskForbiddenWhenRepoScanning.Error())
		// } else if errors.Is(err, service.ErrUserTaskApprovalEnabledError) {
		// 	log.ConsoleW(service.ErrUserTaskApprovalEnabledError.Error())
		// } else {
		zap.L().Error(msg, zap.Error(err))
		os.Exit(1)
		// }
		// 只要有错误都直接退出，区别是退出信号
	}
}

func CheckErr(err error) {
	CheckErrWithMsg(err, "error")
}
