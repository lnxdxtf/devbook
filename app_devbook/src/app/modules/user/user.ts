import { Method, RequestOptions, ResponseJsonSuccess, fetchAPIDevBook } from "../../utils/fetch";
import { NotificationError, NotificationSuccess } from "../../utils/notification";
import { LoginFormDevBookAPI, UserDevBook } from "./interfaces";
/**
 * User class.
 * It implements the UserDevBookAPI interface.
 * User class is used to create to do login and populate the user data if login is successful.
 * This class provides methods to login, logout, password recovery, etc. All About this user.
 * 
 * @implements UserDevBookAPI
 * @class User
 * @module User
 * 
 * @example
 * const user = new User()
 * 
 */
export default class User implements UserDevBook {
    id?: number;
    name?: string;
    nick?: string;
    email?: string;
    token?: string;
    created_at?: string | Date;


    async login(form: LoginFormDevBookAPI): Promise<boolean> {
        try {
            const response = await fetchAPIDevBook({ method: Method.POST, path: '/login', body: form } as RequestOptions) as ResponseJsonSuccess
            this.id = response.data.id
            this.token = response.data.token
            if (this.id) {
                const headers = { 'Authorization': `Bearer ${this.token}` }
                const userDataResponse = await fetchAPIDevBook({ method: Method.GET, path: `/users/${this.id}`, headers: headers } as RequestOptions) as ResponseJsonSuccess
                this.email = userDataResponse.data.email
                this.name = userDataResponse.data.name
                this.created_at = userDataResponse.data.created_at
                this.nick = userDataResponse.data.nick
                NotificationSuccess('Login successful')
                return true
            } else {
                NotificationError('Could not login, please try again later')
                return false
            }
        } catch (error) {
            console.log(error)
            NotificationError('Error attempting to login, please try again later')
            return false
        }



    }

    async getUserPosts() {
        try {
            const headers = { 'Authorization': `Bearer ${this.token}` }
            const response = await fetchAPIDevBook({ method: Method.GET, path: '/posts', headers } as RequestOptions) as ResponseJsonSuccess
            return response.data
        } catch (error) {
            console.log(error)
            NotificationError('Error attempting to get feed, please try again later')
            return []
        }
    }

    public async logout() {
        delete this.id
        delete this.name
        delete this.email
        delete this.token
        delete this.created_at
    }

}