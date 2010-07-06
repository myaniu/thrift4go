/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package server;

import (
 "thrift/transport";
 "thrift/protocol";
 "thrift/base";
 "os";
)

type TServer interface {
  /**
   * Core processor
   */
  ProcessorFactory() base.TProcessorFactory;
  /**
   * Server transport
   */
  ServerTransport() transport.TServerTransport;
  /**
   * Input Transport Factory
   */
  InputTransportFactory() transport.TTransportFactory;
  /**
   * Output Transport Factory
   */
  OutputTransportFactory() transport.TTransportFactory;
  /**
   * Input Protocol Factory
   */
  InputProtocolFactory() protocol.TProtocolFactory;
  /**
   * Output Protocol Factory
   */
  OutputProtocolFactory() protocol.TProtocolFactory;
  
  /**
   * The run method fires up the server and gets things going.
   */
  Serve() os.Error;
  /**
   * Stop the server. This is optional on a per-implementation basis. Not
   * all servers are required to be cleanly stoppable.
   */
  Stop() os.Error;
}




