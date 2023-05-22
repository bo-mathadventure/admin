/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { admin_adminMapChat } from './admin_adminMapChat';

export type admin_createMap = {
    canReport: boolean;
    chat?: admin_adminMapChat;
    contactPage: string;
    expireOn: string;
    mapUrl: string;
    policy: 'anonymous' | 'login' | 'permission';
    roomName: string;
    tags: Array<string>;
};
