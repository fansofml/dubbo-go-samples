/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/client"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	greet "github.com/apache/dubbo-go-samples/java_interop/protobuf-triple/go/proto"
	"github.com/dubbogo/gost/log/logger"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_SAMPLES/rpc/triple/pb/dubbogo-java/go-client/conf/dubbogo.yml
func main() {

	cli, err := client.NewClient(
		client.WithClientURL("127.0.0.1:36969"),
	)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	svc, err := greet.NewGreeter(cli)
	if err != nil {
		panic(err)
	}

	resp, err := svc.SayHello(context.Background(), &greet.HelloRequest{Name: "hello world"})
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("Greet response: %s", resp.Message)
}
