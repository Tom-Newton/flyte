// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts"
// @generated from file flyteidl/service/identity.proto (package flyteidl.service, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { UserInfoRequest, UserInfoResponse } from "./identity_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * IdentityService defines an RPC Service that interacts with user/app identities.
 *
 * @generated from service flyteidl.service.IdentityService
 */
export const IdentityService = {
  typeName: "flyteidl.service.IdentityService",
  methods: {
    /**
     * Retrieves user information about the currently logged in user.
     *
     * @generated from rpc flyteidl.service.IdentityService.UserInfo
     */
    userInfo: {
      name: "UserInfo",
      I: UserInfoRequest,
      O: UserInfoResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

