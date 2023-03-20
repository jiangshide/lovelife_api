package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"],
        beego.ControllerComments{
            Method: "Ads",
            Router: "/ads",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"],
        beego.ControllerComments{
            Method: "Status",
            Router: "/status",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AppController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/update",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AuditBlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AuditBlogController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AuditBlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AuditBlogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AutitChannelController"],
        beego.ControllerComments{
            Method: "Type",
            Router: "/type",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AutitUserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AutitUserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:AutitUserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:AutitUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlockChainController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlockChainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlockChainController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlockChainController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Id",
            Router: "/Id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Browse",
            Router: "/browse",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Channel",
            Router: "/channel",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "City",
            Router: "/city",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Collection",
            Router: "/collection",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "CollectionAdd",
            Router: "/collectionAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Comment",
            Router: "/comment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "CommentAdd",
            Router: "/commentAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "CommentUid",
            Router: "/commentUid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "CommentUidAdd",
            Router: "/commentUidAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Find",
            Router: "/find",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Follow",
            Router: "/follow",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Format",
            Router: "/format",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "HistoryPost",
            Router: "/history",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "HistoryGet",
            Router: "/history",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Lrc",
            Router: "/lrc",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "LrcUodate",
            Router: "/lrcUpdate",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Praise",
            Router: "/praise",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "PraiseAdd",
            Router: "/praiseAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "PraiseCommentAdd",
            Router: "/praiseCommentAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "PraiseCommentUidAdd",
            Router: "/praiseCommentUidAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Publish",
            Router: "/publish",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Recommend",
            Router: "/recommend",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Remove",
            Router: "/remove",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Share",
            Router: "/share",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "ShareAdd",
            Router: "/shareAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "SyncBlog",
            Router: "/syncBlog",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "SyncStatus",
            Router: "/syncStatus",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "TopAdd",
            Router: "/topAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "User",
            Router: "/user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:BlogController"],
        beego.ControllerComments{
            Method: "ViewAdd",
            Router: "/viewAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelId",
            Router: "/channelId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channelType",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Follow",
            Router: "/follow",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Id",
            Router: "/id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Nature",
            Router: "/nature",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Official",
            Router: "/official",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Recommend",
            Router: "/recommend",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Search",
            Router: "/search",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "Type",
            Router: "/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "TypeList",
            Router: "/typeList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ChannelController"],
        beego.ControllerComments{
            Method: "User",
            Router: "/user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "AddInit",
            Router: "/addInit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Banner",
            Router: "/banner",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Del",
            Router: "/del",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Jsd",
            Router: "/jsd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/update",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:FormatController"],
        beego.ControllerComments{
            Method: "Word",
            Router: "/word",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:HomeController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:HomeController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:HomeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:PayController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:PayController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"],
        beego.ControllerComments{
            Method: "Feedback",
            Router: "/feedback",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"],
        beego.ControllerComments{
            Method: "FeedbackList",
            Router: "/feedbackList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:ReportController"],
        beego.ControllerComments{
            Method: "Type",
            Router: "/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UploadController"],
        beego.ControllerComments{
            Method: "Multipart",
            Router: "/multipart",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Active",
            Router: "/active",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "ActiveAdd",
            Router: "/activeAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Bind",
            Router: "/bind",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Certification",
            Router: "/certification",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "CertificationList",
            Router: "/certificationList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "CertificationUpdate",
            Router: "/certificationUpdate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "CodeLogin",
            Router: "/codeLogin",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Course",
            Router: "/course",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Fans",
            Router: "/fans",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Follow",
            Router: "/follow",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "FollowAdd",
            Router: "/followAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Friend",
            Router: "/friend",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "FriendAdd",
            Router: "/friendAdd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "ImUserDelete",
            Router: "/imUserDelete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "ImUserList",
            Router: "/imUserList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Online",
            Router: "/online",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "OnlineList",
            Router: "/onlineList",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Praise",
            Router: "/praise",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Profile",
            Router: "/profile",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Profiles",
            Router: "/profiles",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Reg",
            Router: "/reg",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Remarks",
            Router: "/remarks",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserExit",
            Router: "/userExit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Validate",
            Router: "/validate",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "View",
            Router: "/view",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"] = append(beego.GlobalControllerRouter["sanskrit_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "WeChat",
            Router: "/weChat",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
