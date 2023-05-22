/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { handler_APIResponse } from '../models/handler_APIResponse';
import type { handler_loginRequest } from '../models/handler_loginRequest';
import type { handler_loginResponse } from '../models/handler_loginResponse';
import type { handler_registerRequest } from '../models/handler_registerRequest';
import type { handler_updateUserRequest } from '../models/handler_updateUserRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class AuthService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Get Login Token
     * Do a login with user credentials (email/password) when a password is set
     * @param params -
     * @returns handler_loginResponse OK
     * @throws ApiError
     */
    public postAuthLogin(
params: handler_loginRequest,
): CancelablePromise<handler_loginResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/auth/login',
            body: params,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Register new user
     * Start a registration of a new user. Only works when registration are enabled
     * @param params -
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public postAuthRegister(
params: handler_registerRequest,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/auth/register',
            body: params,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Resend mail confirmation
     * Resend mail confirmation if user not confirmed
     * @param params -
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public postAuthResendConfirmation(
params: handler_updateUserRequest,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/auth/resendConfirmation',
            body: params,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * SAML Response Callback
     * Get SAML response of the IDP. This route is only available when SAML is correctly configured.
     * @param relayState 
     * @param samlResponse 
     * @returns void 
     * @throws ApiError
     */
    public postAuthSamlAcs(
relayState?: string,
samlResponse?: string,
): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/auth/saml/acs',
            formData: {
                'RelayState': relayState,
                'samlResponse': samlResponse,
            },
            errors: {
                302: `Found`,
                400: `Bad Request`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Get SAML Auth URL
     * Starts a new SAML authentication flow. This route is only available when SAML is correctly configured.
     * @returns void 
     * @throws ApiError
     */
    public getAuthSamlStart(): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/auth/saml/start',
            errors: {
                302: `Found`,
                500: `Internal Server Error`,
            },
        });
    }

}
