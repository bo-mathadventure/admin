/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { handler_updateUserRequest } from '../models/handler_updateUserRequest';
import type { handler_userResponse } from '../models/handler_userResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class UserService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * Token
     * execute token actions
     * @param token -
     * @param params -
     * @returns handler_userResponse OK
     * @throws ApiError
     */
    public postAuthToken(
token: string,
params: handler_updateUserRequest,
): CancelablePromise<handler_userResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/auth/token',
            query: {
                'token': token,
            },
            body: params,
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * User Info
     * Get user details of logged-in user
     * @returns handler_userResponse OK
     * @throws ApiError
     */
    public getSystemUser(): CancelablePromise<handler_userResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/user',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update User
     * Update details of the logged-in user
     * @param params -
     * @returns handler_userResponse OK
     * @throws ApiError
     */
    public putSystemUser(
params: handler_updateUserRequest,
): CancelablePromise<handler_userResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/user',
            body: params,
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Workadventure token
     * Generate JWT Token of logged-in user and directly redirect user to Workadventure
     * @returns void 
     * @throws ApiError
     */
    public getSystemUserToken(): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/user/token',
            errors: {
                302: `Found`,
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

}
