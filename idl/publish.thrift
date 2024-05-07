namespace go publishs
enum Code2 {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

struct Video{
    1:i64 VideoId
    2:i64 AuthorId
    3:string PlayUrl
    4:string CoverUrl
    5:i64 FavoriteCount
    6:i64 CommentCount
    7:string PublishTime
    8:string Title
}

struct UpLoadVideoRequest{
    1: string ContentType     (api.body="content_type", api.form="content_type",api.vd="(len($) > 0 && len($) < 100)")
    2: string ObjectName    (api.body="object_name", api.form="object_name",api.vd="(len($) > 0 && len($) < 100)")
    3: string BucketName    (api.body="bucket_name", api.form="bucket_name",api.vd="(len($) > 0 && len($) < 100)")
    4: i64 UserId
    5: string path
    6: string Title
    7: string CoverUrl
}

struct UpLoadVideoResponse{
   1: Code2 code
   2: string msg
}

struct VideoCreateRequest{
    1:Video video
}
struct VideoCreateResponse{
    1:Code2 code
    2:string msg
}

service UpLoadVideoService {
   VideoCreateResponse VideoCreate (1:VideoCreateRequest req)
   UpLoadVideoResponse UploadVideo(1:UpLoadVideoRequest req)(api.post="/v1/video/upload")
}