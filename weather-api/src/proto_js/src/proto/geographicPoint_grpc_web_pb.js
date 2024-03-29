/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!

/* eslint-disable */
// @ts-nocheck

const grpc = {};
grpc.web = require("grpc-web");

const proto = {};
proto.pb = require("./geographicPoint_pb.js");

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.GeographicPointServiceClient = function (
  hostname,
  credentials,
  options
) {
  if (!options) options = {};
  options.format = "text";

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.GeographicPointServicePromiseClient = function (
  hostname,
  credentials,
  options
) {
  if (!options) options = {};
  options.format = "text";

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;
};

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.AddGeographicPointRequest,
 *   !proto.pb.RegsiterGeographicPointResponse>}
 */
const methodDescriptor_GeographicPointService_AddGeographicPoint =
  new grpc.web.MethodDescriptor(
    "/pb.GeographicPointService/AddGeographicPoint",
    grpc.web.MethodType.UNARY,
    proto.pb.AddGeographicPointRequest,
    proto.pb.RegsiterGeographicPointResponse,
    /**
     * @param {!proto.pb.AddGeographicPointRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.RegsiterGeographicPointResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.AddGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.RegsiterGeographicPointResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.RegsiterGeographicPointResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.GeographicPointServiceClient.prototype.addGeographicPoint = function (
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + "/pb.GeographicPointService/AddGeographicPoint",
    request,
    metadata || {},
    methodDescriptor_GeographicPointService_AddGeographicPoint,
    callback
  );
};

/**
 * @param {!proto.pb.AddGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.RegsiterGeographicPointResponse>}
 *     Promise that resolves to the response
 */
proto.pb.GeographicPointServicePromiseClient.prototype.addGeographicPoint =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ + "/pb.GeographicPointService/AddGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_AddGeographicPoint
    );
  };

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.UpdateGeographicPointRequest,
 *   !proto.pb.RegsiterGeographicPointResponse>}
 */
const methodDescriptor_GeographicPointService_UpdateGeographicPoint =
  new grpc.web.MethodDescriptor(
    "/pb.GeographicPointService/UpdateGeographicPoint",
    grpc.web.MethodType.UNARY,
    proto.pb.UpdateGeographicPointRequest,
    proto.pb.RegsiterGeographicPointResponse,
    /**
     * @param {!proto.pb.UpdateGeographicPointRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.RegsiterGeographicPointResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.UpdateGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.RegsiterGeographicPointResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.RegsiterGeographicPointResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.GeographicPointServiceClient.prototype.updateGeographicPoint =
  function (request, metadata, callback) {
    return this.client_.rpcCall(
      this.hostname_ + "/pb.GeographicPointService/UpdateGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_UpdateGeographicPoint,
      callback
    );
  };

/**
 * @param {!proto.pb.UpdateGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.RegsiterGeographicPointResponse>}
 *     Promise that resolves to the response
 */
proto.pb.GeographicPointServicePromiseClient.prototype.updateGeographicPoint =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ + "/pb.GeographicPointService/UpdateGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_UpdateGeographicPoint
    );
  };

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.DeleteGeographicPointRequest,
 *   !proto.pb.RegsiterGeographicPointResponse>}
 */
const methodDescriptor_GeographicPointService_DeleteGeographicPoint =
  new grpc.web.MethodDescriptor(
    "/pb.GeographicPointService/DeleteGeographicPoint",
    grpc.web.MethodType.UNARY,
    proto.pb.DeleteGeographicPointRequest,
    proto.pb.RegsiterGeographicPointResponse,
    /**
     * @param {!proto.pb.DeleteGeographicPointRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.RegsiterGeographicPointResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.DeleteGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.RegsiterGeographicPointResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.RegsiterGeographicPointResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.GeographicPointServiceClient.prototype.deleteGeographicPoint =
  function (request, metadata, callback) {
    return this.client_.rpcCall(
      this.hostname_ + "/pb.GeographicPointService/DeleteGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_DeleteGeographicPoint,
      callback
    );
  };

/**
 * @param {!proto.pb.DeleteGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.RegsiterGeographicPointResponse>}
 *     Promise that resolves to the response
 */
proto.pb.GeographicPointServicePromiseClient.prototype.deleteGeographicPoint =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ + "/pb.GeographicPointService/DeleteGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_DeleteGeographicPoint
    );
  };

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.GetWeatherListByGeographicPointRequest,
 *   !proto.pb.GetWeatherListByGeographicPointResponse>}
 */
const methodDescriptor_GeographicPointService_GetWeatherListByGeographicPoint =
  new grpc.web.MethodDescriptor(
    "/pb.GeographicPointService/GetWeatherListByGeographicPoint",
    grpc.web.MethodType.UNARY,
    proto.pb.GetWeatherListByGeographicPointRequest,
    proto.pb.GetWeatherListByGeographicPointResponse,
    /**
     * @param {!proto.pb.GetWeatherListByGeographicPointRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.GetWeatherListByGeographicPointResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.GetWeatherListByGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.GetWeatherListByGeographicPointResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.GetWeatherListByGeographicPointResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.GeographicPointServiceClient.prototype.getWeatherListByGeographicPoint =
  function (request, metadata, callback) {
    return this.client_.rpcCall(
      this.hostname_ +
        "/pb.GeographicPointService/GetWeatherListByGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_GetWeatherListByGeographicPoint,
      callback
    );
  };

/**
 * @param {!proto.pb.GetWeatherListByGeographicPointRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.GetWeatherListByGeographicPointResponse>}
 *     Promise that resolves to the response
 */
proto.pb.GeographicPointServicePromiseClient.prototype.getWeatherListByGeographicPoint =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ +
        "/pb.GeographicPointService/GetWeatherListByGeographicPoint",
      request,
      metadata || {},
      methodDescriptor_GeographicPointService_GetWeatherListByGeographicPoint
    );
  };

module.exports = proto.pb;
