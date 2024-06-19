/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { IdentityData } from "./IdentityData";
/**
 * Authentication data set
 */
export type AuthSetDeviceauth = {
  /**
   * Authentication data set ID.
   */
  id?: string;
  /**
   * The device's public key (PEM encoding), generated by the device or pre-provisioned by the vendor. Currently supported public algorithms are: RSA, ED25519 and ECDSA.
   */
  pubkey?: string;
  identity_data?: IdentityData;
  status?: AuthSetDeviceauth.status;
  /**
   * Created timestamp
   */
  ts?: string;
  /**
   * Device ID connected to authentication data set
   */
  device_id?: string;
};
export namespace AuthSetDeviceauth {
  export enum status {
    PENDING = "pending",
    ACCEPTED = "accepted",
    REJECTED = "rejected",
    PREAUTHORIZED = "preauthorized",
    NOAUTH = "noauth",
  }
}
