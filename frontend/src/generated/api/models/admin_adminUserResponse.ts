/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { admin_adminGroupResponse } from './admin_adminGroupResponse';

export type admin_adminUserResponse = {
    createdAt?: string;
    email?: string;
    groups?: Array<admin_adminGroupResponse>;
    id?: number;
    lastLogin?: string;
    permissions?: Array<string>;
    ssoIdentifier?: string;
    tags?: Array<string>;
    userPermissions?: Array<string>;
    userTags?: Array<string>;
    username?: string;
    uuid?: string;
};
