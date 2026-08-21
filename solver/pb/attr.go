package pb

const (
	AttrKeepGitDir        = "git.keepgitdir"
	AttrFullRemoteURL     = "git.fullurl"
	AttrAuthHeaderSecret  = "git.authheadersecret"
	AttrAuthTokenSecret   = "git.authtokensecret"
	AttrKnownSSHHosts     = "git.knownsshhosts"
	AttrMountSSHSock      = "git.mountsshsock"
	AttrGitChecksum       = "git.checksum"
	AttrGitSkipSubmodules = "git.skipsubmodules"
	AttrGitMTime          = "git.mtime"
	AttrGitFetchByCommit  = "git.fetchbycommit"
	AttrGitBundle         = "git.bundle"
	AttrGitCheckoutBundle = "git.checkoutbundle"
)

const (
	AttrGitSignatureVerifyPubKey           = "git.sig.pubkey"
	AttrGitSignatureVerifyRejectExpired    = "git.sig.rejectexpired"
	AttrGitSignatureVerifyRequireSignedTag = "git.sig.requiresignedtag"
	AttrGitSignatureVerifyIgnoreSignedTag  = "git.sig.ignoresignedtag"
)

const (
	AttrLocalSessionID          = "local.session"
	AttrLocalUniqueID           = "local.unique"
	AttrIncludePatterns         = "local.includepattern"
	AttrFollowPaths             = "local.followpaths"
	AttrExcludePatterns         = "local.excludepatterns"
	AttrSharedKeyHint           = "local.sharedkeyhint"
	AttrMetadataTransfer        = "local.metadatatransfer"
	AttrMetadataTransferExclude = "local.metadatatransferexclude"
)

const AttrLLBDefinitionFilename = "llbbuild.filename"

const (
	AttrHTTPChecksum              = "http.checksum"
	AttrHTTPFilename              = "http.filename"
	AttrHTTPPerm                  = "http.perm"
	AttrHTTPUID                   = "http.uid"
	AttrHTTPGID                   = "http.gid"
	AttrHTTPAuthHeaderSecret      = "http.authheadersecret"
	AttrHTTPHeaderPrefix          = "http.header."
	AttrHTTPSignatureVerifyPubKey = "http.sig.pubkey"
	AttrHTTPSignatureVerify       = "http.sig.signature"
)

const (
	AttrImageResolveMode            = "image.resolvemode"
	AttrImageResolveModeDefault     = "default"
	AttrImageResolveModeForcePull   = "pull"
	AttrImageResolveModePreferLocal = "local"
	AttrImageRecordType             = "image.recordtype"
	AttrImageLayerLimit             = "image.layerlimit"
	AttrImageChecksum               = "image.checksum"
)

const (
	AttrOCILayoutSessionID  = "oci.session"
	AttrOCILayoutStoreID    = "oci.store"
	AttrOCILayoutLayerLimit = "oci.layerlimit"
)

const (
	AttrLocalDiffer         = "local.differ"
	AttrLocalDifferNone     = "none"
	AttrLocalDifferMetadata = "metadata"
)

type IsFileAction = isFileAction_Action
