/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { config_Config } from './config_Config';

export type handler_userResponse = {
    config?: config_Config;
    createdAt?: string;
    email?: string;
    lastLogin?: string;
    permissions?: Array<string>;
    tags?: Array<string>;
    username?: string;
    uuid?: string;
};
