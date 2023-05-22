/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class SamlService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

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
