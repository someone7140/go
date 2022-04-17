// source: src/proto/geographicPoint.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require("google-protobuf");
var goog = jspb;
var global = function () {
  if (this) {
    return this;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  return Function("return this")();
}.call(null);

goog.exportSymbol("proto.pb.AddGeographicPointRequest", null, global);
goog.exportSymbol("proto.pb.RegsiterGeographicPointResponse", null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.AddGeographicPointRequest = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.AddGeographicPointRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.AddGeographicPointRequest.displayName =
    "proto.pb.AddGeographicPointRequest";
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.RegsiterGeographicPointResponse = function (opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.RegsiterGeographicPointResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.RegsiterGeographicPointResponse.displayName =
    "proto.pb.RegsiterGeographicPointResponse";
}

if (jspb.Message.GENERATE_TO_OBJECT) {
  /**
   * Creates an object representation of this proto.
   * Field names that are reserved in JavaScript and will be renamed to pb_name.
   * Optional fields that are not set will be set to undefined.
   * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
   * For the list of reserved names please see:
   *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
   * @param {boolean=} opt_includeInstance Deprecated. whether to include the
   *     JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @return {!Object}
   */
  proto.pb.AddGeographicPointRequest.prototype.toObject = function (
    opt_includeInstance
  ) {
    return proto.pb.AddGeographicPointRequest.toObject(
      opt_includeInstance,
      this
    );
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.pb.AddGeographicPointRequest} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.pb.AddGeographicPointRequest.toObject = function (
    includeInstance,
    msg
  ) {
    var f,
      obj = {
        name: jspb.Message.getFieldWithDefault(msg, 1, ""),
        lat: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
        lon: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
      };

    if (includeInstance) {
      obj.$jspbMessageInstance = msg;
    }
    return obj;
  };
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.AddGeographicPointRequest}
 */
proto.pb.AddGeographicPointRequest.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.AddGeographicPointRequest();
  return proto.pb.AddGeographicPointRequest.deserializeBinaryFromReader(
    msg,
    reader
  );
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.AddGeographicPointRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.AddGeographicPointRequest}
 */
proto.pb.AddGeographicPointRequest.deserializeBinaryFromReader = function (
  msg,
  reader
) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
      case 1:
        var value = /** @type {string} */ (reader.readString());
        msg.setName(value);
        break;
      case 2:
        var value = /** @type {number} */ (reader.readDouble());
        msg.setLat(value);
        break;
      case 3:
        var value = /** @type {number} */ (reader.readDouble());
        msg.setLon(value);
        break;
      default:
        reader.skipField();
        break;
    }
  }
  return msg;
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.AddGeographicPointRequest.prototype.serializeBinary = function () {
  var writer = new jspb.BinaryWriter();
  proto.pb.AddGeographicPointRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.AddGeographicPointRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.AddGeographicPointRequest.serializeBinaryToWriter = function (
  message,
  writer
) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(1, f);
  }
  f = message.getLat();
  if (f !== 0.0) {
    writer.writeDouble(2, f);
  }
  f = message.getLon();
  if (f !== 0.0) {
    writer.writeDouble(3, f);
  }
};

/**
 * optional string name = 1;
 * @return {string}
 */
proto.pb.AddGeographicPointRequest.prototype.getName = function () {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};

/**
 * @param {string} value
 * @return {!proto.pb.AddGeographicPointRequest} returns this
 */
proto.pb.AddGeographicPointRequest.prototype.setName = function (value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};

/**
 * optional double lat = 2;
 * @return {number}
 */
proto.pb.AddGeographicPointRequest.prototype.getLat = function () {
  return /** @type {number} */ (
    jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0)
  );
};

/**
 * @param {number} value
 * @return {!proto.pb.AddGeographicPointRequest} returns this
 */
proto.pb.AddGeographicPointRequest.prototype.setLat = function (value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};

/**
 * optional double lon = 3;
 * @return {number}
 */
proto.pb.AddGeographicPointRequest.prototype.getLon = function () {
  return /** @type {number} */ (
    jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0)
  );
};

/**
 * @param {number} value
 * @return {!proto.pb.AddGeographicPointRequest} returns this
 */
proto.pb.AddGeographicPointRequest.prototype.setLon = function (value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};

if (jspb.Message.GENERATE_TO_OBJECT) {
  /**
   * Creates an object representation of this proto.
   * Field names that are reserved in JavaScript and will be renamed to pb_name.
   * Optional fields that are not set will be set to undefined.
   * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
   * For the list of reserved names please see:
   *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
   * @param {boolean=} opt_includeInstance Deprecated. whether to include the
   *     JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @return {!Object}
   */
  proto.pb.RegsiterGeographicPointResponse.prototype.toObject = function (
    opt_includeInstance
  ) {
    return proto.pb.RegsiterGeographicPointResponse.toObject(
      opt_includeInstance,
      this
    );
  };

  /**
   * Static version of the {@see toObject} method.
   * @param {boolean|undefined} includeInstance Deprecated. Whether to include
   *     the JSPB instance for transitional soy proto support:
   *     http://goto/soy-param-migration
   * @param {!proto.pb.RegsiterGeographicPointResponse} msg The msg instance to transform.
   * @return {!Object}
   * @suppress {unusedLocalVariables} f is only used for nested messages
   */
  proto.pb.RegsiterGeographicPointResponse.toObject = function (
    includeInstance,
    msg
  ) {
    var f,
      obj = {};

    if (includeInstance) {
      obj.$jspbMessageInstance = msg;
    }
    return obj;
  };
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.RegsiterGeographicPointResponse}
 */
proto.pb.RegsiterGeographicPointResponse.deserializeBinary = function (bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.RegsiterGeographicPointResponse();
  return proto.pb.RegsiterGeographicPointResponse.deserializeBinaryFromReader(
    msg,
    reader
  );
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.RegsiterGeographicPointResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.RegsiterGeographicPointResponse}
 */
proto.pb.RegsiterGeographicPointResponse.deserializeBinaryFromReader =
  function (msg, reader) {
    while (reader.nextField()) {
      if (reader.isEndGroup()) {
        break;
      }
      var field = reader.getFieldNumber();
      switch (field) {
        default:
          reader.skipField();
          break;
      }
    }
    return msg;
  };

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.RegsiterGeographicPointResponse.prototype.serializeBinary =
  function () {
    var writer = new jspb.BinaryWriter();
    proto.pb.RegsiterGeographicPointResponse.serializeBinaryToWriter(
      this,
      writer
    );
    return writer.getResultBuffer();
  };

/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.RegsiterGeographicPointResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.RegsiterGeographicPointResponse.serializeBinaryToWriter = function (
  message,
  writer
) {
  var f = undefined;
};

goog.object.extend(exports, proto.pb);
