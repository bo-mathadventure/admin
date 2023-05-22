/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { admin_adminMapChat } from './admin_adminMapChat';

export type admin_adminMapResponse = {
    canReport?: boolean;
    chat?: admin_adminMapChat;
    contactPage?: string;
    createdAt?: string;
    expireOn?: string;
    id?: number;
    mapUrl?: string;
    policy?: 'anonymous' | 'login' | 'permission';
    roomName?: string;
    tags?: Array<string>;
};
