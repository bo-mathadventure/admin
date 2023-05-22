/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { handler_ErrorResponse } from './handler_ErrorResponse';

export type handler_APIResponse = {
    error?: boolean;
    extra?: string;
    message?: string;
    status?: number;
    success?: boolean;
    validation?: Array<handler_ErrorResponse>;
};
