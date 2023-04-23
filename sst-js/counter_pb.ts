// @generated by protoc-gen-es v1.2.0 with parameter "target=ts,import_extension=.ts"
// @generated from file counter.proto (package sst, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message sst.CountToRequest
 */
export class CountToRequest extends Message<CountToRequest> {
  /**
   * The number to count to.
   *
   * @generated from field: int32 to = 1;
   */
  to = 0;

  constructor(data?: PartialMessage<CountToRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "sst.CountToRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "to", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CountToRequest {
    return new CountToRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CountToRequest {
    return new CountToRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CountToRequest {
    return new CountToRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CountToRequest | PlainMessage<CountToRequest> | undefined, b: CountToRequest | PlainMessage<CountToRequest> | undefined): boolean {
    return proto3.util.equals(CountToRequest, a, b);
  }
}

/**
 * @generated from message sst.CountToResponse
 */
export class CountToResponse extends Message<CountToResponse> {
  /**
   * The current count.
   *
   * @generated from field: int32 count = 1;
   */
  count = 0;

  constructor(data?: PartialMessage<CountToResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "sst.CountToResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CountToResponse {
    return new CountToResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CountToResponse {
    return new CountToResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CountToResponse {
    return new CountToResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CountToResponse | PlainMessage<CountToResponse> | undefined, b: CountToResponse | PlainMessage<CountToResponse> | undefined): boolean {
    return proto3.util.equals(CountToResponse, a, b);
  }
}

