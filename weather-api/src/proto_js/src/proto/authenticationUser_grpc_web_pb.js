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
proto.pb = require("./authenticationUser_pb.js");

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.AuthenticationUserServiceClient = function (
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
proto.pb.AuthenticationUserServicePromiseClient = function (
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
 *   !proto.pb.VerifyGoogleAuthCodeRequest,
 *   !proto.pb.UserResponse>}
 */
const methodDescriptor_AuthenticationUserService_VerifyGoogleAuthCode =
  new grpc.web.MethodDescriptor(
    "/pb.AuthenticationUserService/VerifyGoogleAuthCode",
    grpc.web.MethodType.UNARY,
    proto.pb.VerifyGoogleAuthCodeRequest,
    proto.pb.UserResponse,
    /**
     * @param {!proto.pb.VerifyGoogleAuthCodeRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.UserResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.VerifyGoogleAuthCodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.UserResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.UserResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.AuthenticationUserServiceClient.prototype.verifyGoogleAuthCode =
  function (request, metadata, callback) {
    return this.client_.rpcCall(
      this.hostname_ + "/pb.AuthenticationUserService/VerifyGoogleAuthCode",
      request,
      metadata || {},
      methodDescriptor_AuthenticationUserService_VerifyGoogleAuthCode,
      callback
    );
  };

/**
 * @param {!proto.pb.VerifyGoogleAuthCodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.UserResponse>}
 *     Promise that resolves to the response
 */
proto.pb.AuthenticationUserServicePromiseClient.prototype.verifyGoogleAuthCode =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ + "/pb.AuthenticationUserService/VerifyGoogleAuthCode",
      request,
      metadata || {},
      methodDescriptor_AuthenticationUserService_VerifyGoogleAuthCode
    );
  };

/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.VerifyAuthTokenRequest,
 *   !proto.pb.AuthTokenResponse>}
 */
const methodDescriptor_AuthenticationUserService_VerifyAuthToken =
  new grpc.web.MethodDescriptor(
    "/pb.AuthenticationUserService/VerifyAuthToken",
    grpc.web.MethodType.UNARY,
    proto.pb.VerifyAuthTokenRequest,
    proto.pb.AuthTokenResponse,
    /**
     * @param {!proto.pb.VerifyAuthTokenRequest} request
     * @return {!Uint8Array}
     */
    function (request) {
      return request.serializeBinary();
    },
    proto.pb.AuthTokenResponse.deserializeBinary
  );

/**
 * @param {!proto.pb.VerifyAuthTokenRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pb.AuthTokenResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.AuthTokenResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.AuthenticationUserServiceClient.prototype.verifyAuthToken = function (
  request,
  metadata,
  callback
) {
  return this.client_.rpcCall(
    this.hostname_ + "/pb.AuthenticationUserService/VerifyAuthToken",
    request,
    metadata || {},
    methodDescriptor_AuthenticationUserService_VerifyAuthToken,
    callback
  );
};

/**
 * @param {!proto.pb.VerifyAuthTokenRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.AuthTokenResponse>}
 *     Promise that resolves to the response
 */
proto.pb.AuthenticationUserServicePromiseClient.prototype.verifyAuthToken =
  function (request, metadata) {
    return this.client_.unaryCall(
      this.hostname_ + "/pb.AuthenticationUserService/VerifyAuthToken",
      request,
      metadata || {},
      methodDescriptor_AuthenticationUserService_VerifyAuthToken
    );
  };

module.exports = proto.pb;
