# Mattermost boobs lambda
1. Build: `GOOS=linux go build -o main`
2. Archive: `rm -rf deployment.zip && zip deployment.zip main && rm -rf main`
3. Upload...
3.1. Via aws cli: `aws lambda update-function-code --profile self-mike --function-name mattermost-boobs --zip-file fileb://deployment.zip`

# Mapping template for API Gateway:
`
{
    #foreach( $token in $input.path('$').split('&') )
        #set( $keyVal = $token.split('=') )
        #set( $keyValSize = $keyVal.size() )
        #if( $keyValSize >= 1 )
            #set( $key = $util.urlDecode($keyVal[0]) )
            #if( $keyValSize >= 2 )
                #set( $val = $util.urlDecode($keyVal[1]) )
            #else
                #set( $val = '' )
            #end
            "$key": "$val"#if($foreach.hasNext),#end
        #end
    #end
}
`