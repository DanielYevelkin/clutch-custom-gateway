module.exports = {
    // "@clutch-sh/ec2": {
    //     terminateInstance: {
    //         trending: true,
    //         componentProps: {
    //             resolverType: "clutch.aws.ec2.v1.Instance",
    //         },
    //     },
    //     resizeAutoscalingGroup: {
    //         trending: true,
    //         componentProps: {
    //             resolverType: "clutch.aws.ec2.v1.AutoscalingGroup",
    //         },
    //     },
    // },
    // "@clutch-sh/envoy": {
    //     remoteTriage: {
    //         trending: true,
    //         componentProps: {
    //             options: {
    //                 "Clusters": "clusters",
    //                 "Listeners": "listeners",
    //                 "Runtime": "runtime",
    //                 "Stats": "stats",
    //                 "Server Info": "serverInfo",
    //             },
    //         },
    //     },
    // },
    "@clutch-sh/amiibo": {
        landing: {
          trending: true,
        },
    },
    "@clutch-custom-gateway/echo": {
        echo: {
            trending: true,
        }
    }
};
