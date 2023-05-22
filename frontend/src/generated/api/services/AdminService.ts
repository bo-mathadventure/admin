/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { admin_adminAnnouncementResponse } from '../models/admin_adminAnnouncementResponse';
import type { admin_adminBanResponse } from '../models/admin_adminBanResponse';
import type { admin_adminGroupResponse } from '../models/admin_adminGroupResponse';
import type { admin_adminMapResponse } from '../models/admin_adminMapResponse';
import type { admin_adminReportResponse } from '../models/admin_adminReportResponse';
import type { admin_adminTextureResponse } from '../models/admin_adminTextureResponse';
import type { admin_adminUserResponse } from '../models/admin_adminUserResponse';
import type { admin_createAnnouncement } from '../models/admin_createAnnouncement';
import type { admin_createBan } from '../models/admin_createBan';
import type { admin_createGroup } from '../models/admin_createGroup';
import type { admin_createMap } from '../models/admin_createMap';
import type { admin_createUser } from '../models/admin_createUser';
import type { admin_updateAnnouncement } from '../models/admin_updateAnnouncement';
import type { admin_updateGroup } from '../models/admin_updateGroup';
import type { admin_updateMap } from '../models/admin_updateMap';
import type { admin_updateReport } from '../models/admin_updateReport';
import type { admin_updateUser } from '../models/admin_updateUser';
import type { handler_APIResponse } from '../models/handler_APIResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class AdminService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * List announcements
     * Get all announcements. Requires permission admin.announcement.view or admin.announcement.edit
     * @returns admin_adminAnnouncementResponse OK
     * @throws ApiError
     */
    public getSystemAdminAnnouncement(): CancelablePromise<Array<admin_adminAnnouncementResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/announcement',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create a new announcement
     * Requires permission admin.announcement.edit
     * @param params -
     * @returns admin_adminAnnouncementResponse OK
     * @throws ApiError
     */
    public postSystemAdminAnnouncement(
params: admin_createAnnouncement,
): CancelablePromise<admin_adminAnnouncementResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/announcement',
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
     * Get specific announcement
     * Get announcement by ID. Requires permission admin.announcement.view or admin.announcement.edit
     * @param id Announcement ID
     * @returns admin_adminAnnouncementResponse OK
     * @throws ApiError
     */
    public getSystemAdminAnnouncement1(
id: number,
): CancelablePromise<admin_adminAnnouncementResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/announcement/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update announcement
     * Update announcement by ID. Requires Permission admin.announcement.edit
     * @param params -
     * @param id Announcement ID
     * @returns admin_adminAnnouncementResponse OK
     * @throws ApiError
     */
    public putSystemAdminAnnouncement(
params: admin_updateAnnouncement,
id: number,
): CancelablePromise<admin_adminAnnouncementResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/announcement/{id}',
            path: {
                'id': id,
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
     * Delete announcement
     * Delete announcement by ID. Requires permission admin.announcement.edit
     * @param id Announcement ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminAnnouncement(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/announcement/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List bans
     * Get all bans. Requires permission admin.ban.view or admin.ban.create or admin.ban.delete
     * @returns admin_adminBanResponse OK
     * @throws ApiError
     */
    public getSystemAdminBan(): CancelablePromise<Array<admin_adminBanResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/ban',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create ban
     * Create a ban for an identifier or ip address. Requires permission admin.ban.create
     * @param params -
     * @returns admin_adminBanResponse OK
     * @throws ApiError
     */
    public postSystemAdminBan(
params: admin_createBan,
): CancelablePromise<admin_adminBanResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/ban',
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
     * Get specific ban
     * Get ban by ID. Requires permission admin.ban.view or admin.ban.create or admin.ban.delete
     * @param id Ban ID
     * @returns admin_adminBanResponse OK
     * @throws ApiError
     */
    public getSystemAdminBan1(
id: number,
): CancelablePromise<admin_adminBanResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/ban/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Delete ban
     * Delete ban by ID. Sets validUntil to now() to keep history. Requires permission admin.ban.delete
     * @param id Ban ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminBan(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/ban/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List Groups
     * Get all groups. Requires permission admin.group.edit
     * @returns admin_adminGroupResponse OK
     * @throws ApiError
     */
    public getSystemAdminGroup(): CancelablePromise<Array<admin_adminGroupResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/group',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create Group
     * Create new Group. Requires permission admin.group.edit
     * @param params -
     * @returns admin_adminGroupResponse OK
     * @throws ApiError
     */
    public postSystemAdminGroup(
params: admin_createGroup,
): CancelablePromise<admin_adminGroupResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/group',
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
     * Get specific group
     * Get group by ID. Requires permission admin.group.edit
     * @param id Group ID
     * @returns admin_adminGroupResponse OK
     * @throws ApiError
     */
    public getSystemAdminGroup1(
id: number,
): CancelablePromise<admin_adminGroupResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/group/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update group
     * Update group by ID. Requires permission admin.group.edit
     * @param params -
     * @param id Group ID
     * @returns admin_adminGroupResponse OK
     * @throws ApiError
     */
    public putSystemAdminGroup(
params: admin_updateGroup,
id: number,
): CancelablePromise<admin_adminGroupResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/group/{id}',
            path: {
                'id': id,
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
     * Delete group
     * Delete group by ID. Requires permission admin.group.edit
     * @param id Group ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminGroup(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/group/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List maps
     * Get all maps. Requires permission admin.map.view or admin.map.edit
     * @returns admin_adminMapResponse OK
     * @throws ApiError
     */
    public getSystemAdminMap(): CancelablePromise<Array<admin_adminMapResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/map',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create map
     * Requires permission admin.map.edit
     * @param params -
     * @returns admin_adminMapResponse OK
     * @throws ApiError
     */
    public postSystemAdminMap(
params: admin_createMap,
): CancelablePromise<admin_adminMapResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/map',
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
     * Get specific map
     * Get map by ID. Requires permission admin.map.view or admin.map.edit
     * @param id Map ID
     * @returns admin_adminMapResponse OK
     * @throws ApiError
     */
    public getSystemAdminMap1(
id: number,
): CancelablePromise<admin_adminMapResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/map/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update map
     * Update map by ID. Requires permission admin.map.edit
     * @param params -
     * @param id Map ID
     * @returns admin_adminMapResponse OK
     * @throws ApiError
     */
    public putSystemAdminMap(
params: admin_updateMap,
id: number,
): CancelablePromise<admin_adminMapResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/map/{id}',
            path: {
                'id': id,
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
     * Delete map
     * Delete map by ID. Requires permission admin.map.edit
     * @param id Map ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminMap(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/map/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List reports
     * Get all reports. Requires permission admin.report.view or admin.report.edit
     * @returns admin_adminReportResponse OK
     * @throws ApiError
     */
    public getSystemAdminReport(): CancelablePromise<Array<admin_adminReportResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/report',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Get specific report
     * Get report by ID. Requires permission admin.report.view pr admin.report.edit
     * @param id Report ID
     * @returns admin_adminReportResponse OK
     * @throws ApiError
     */
    public getSystemAdminReport1(
id: number,
): CancelablePromise<admin_adminReportResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/report/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update report
     * Update report by ID. Requires permission admin.report.edit
     * @param params -
     * @param id Report ID
     * @returns admin_adminReportResponse OK
     * @throws ApiError
     */
    public putSystemAdminReport(
params: admin_updateReport,
id: number,
): CancelablePromise<admin_adminReportResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/report/{id}',
            path: {
                'id': id,
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
     * Delete report
     * Delete report by ID. Sets hide to true to keep history. Requires permission admin.report.edit
     * @param id Report ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminReport(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/report/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List textures
     * Get all textures. Requires permission admin.texture.view or admin.texture.edit
     * @returns admin_adminTextureResponse OK
     * @throws ApiError
     */
    public getSystemAdminTexture(): CancelablePromise<Array<admin_adminTextureResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/texture',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create texture
     * Upload file via resource field. Requires permission admin.texture.edit
     * @param layer 
     * @param texture 
     * @param resource the texture file
     * @param tags 
     * @returns admin_adminTextureResponse OK
     * @throws ApiError
     */
    public postSystemAdminTexture(
layer: 'woka' | 'body' | 'hair' | 'eyes' | 'hat' | 'accessory' | 'clothes' | 'companion',
texture: string,
resource: Blob,
tags?: Array<string>,
): CancelablePromise<admin_adminTextureResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/texture',
            formData: {
                'layer': layer,
                'tags': tags,
                'texture': texture,
                'resource': resource,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Get specific texture
     * Get texture by ID. Requires permission admin.texture.view or admin.texture.edit
     * @param id Texture ID
     * @returns admin_adminTextureResponse OK
     * @throws ApiError
     */
    public getSystemAdminTexture1(
id: number,
): CancelablePromise<admin_adminTextureResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/texture/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update texture
     * Update texture by ID. Requires permission admin.texture.edit
     * @param layer 
     * @param texture 
     * @param id Texture ID
     * @param resource the texture file
     * @param tags 
     * @returns admin_adminTextureResponse OK
     * @throws ApiError
     */
    public putSystemAdminTexture(
layer: 'woka' | 'body' | 'hair' | 'eyes' | 'hat' | 'accessory' | 'clothes' | 'companion',
texture: string,
id: number,
resource: Blob,
tags?: Array<string>,
): CancelablePromise<admin_adminTextureResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/texture/{id}',
            path: {
                'id': id,
            },
            formData: {
                'layer': layer,
                'tags': tags,
                'texture': texture,
                'resource': resource,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Delete texture
     * Delete texture by ID. Requires permission admin.texture.edit
     * @param id Texture ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminTexture(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/texture/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * List users
     * Get all users. Requires permission admin.user.view or admin.user.edit
     * @returns admin_adminUserResponse OK
     * @throws ApiError
     */
    public getSystemAdminUser(): CancelablePromise<Array<admin_adminUserResponse>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/user',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Invite user
     * Invite/Create new user. Requires permission admin.user.invite
     * @param params -
     * @returns admin_adminUserResponse OK
     * @throws ApiError
     */
    public postSystemAdminUserInvite(
params: admin_createUser,
): CancelablePromise<admin_adminUserResponse> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/system/admin/user/invite',
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
     * Get specific user
     * Get user by ID. Requires permission admin.user.view or admin.user.edit
     * @param id User ID
     * @returns admin_adminUserResponse OK
     * @throws ApiError
     */
    public getSystemAdminUser1(
id: number,
): CancelablePromise<admin_adminUserResponse> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/system/admin/user/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update user
     * Update user by ID. Requires permission admin.user.edit
     * @param params -
     * @param id User ID
     * @returns admin_adminUserResponse OK
     * @throws ApiError
     */
    public putSystemAdminUser(
params: admin_updateUser,
id: number,
): CancelablePromise<admin_adminUserResponse> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/system/admin/user/{id}',
            path: {
                'id': id,
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
     * Delete user
     * Delete user by ID. Requires permission admin.user.edit
     * @param id User ID
     * @returns handler_APIResponse OK
     * @throws ApiError
     */
    public deleteSystemAdminUser(
id: number,
): CancelablePromise<handler_APIResponse> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/system/admin/user/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

}
