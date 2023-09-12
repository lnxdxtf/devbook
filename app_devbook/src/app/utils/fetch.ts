import { NotificationError } from "./notification"

export interface ResponseJsonSuccess {
    data: any
}

export interface ResponseJsonError {
    error?: string | any,
    errors?: string[] | any | any[]
}


export enum Method {
    GET = 'GET',
    POST = 'POST',
    PUT = 'PUT',
    DELETE = 'DELETE'
}


export interface RequestOptions {
    path: string,
    method: Method | string,
    headers?: Headers | any,
    body?: BodyInit | Object
}

export async function fetchAPIDevBook(req_opt: RequestOptions): Promise<ResponseJsonSuccess | ResponseJsonError> {
    let baseHeaders: Headers | any
    if (req_opt.headers) {
        baseHeaders = { 'Content-Type': 'application/json' }
        Object.assign(baseHeaders, req_opt.headers)
    } else {
        baseHeaders = new Headers()
        baseHeaders.append('Content-Type', 'application/json')
    }

    const options = {
        method: req_opt.method,
        headers: baseHeaders,
        body: req_opt.body ? JSON.stringify(req_opt.body) : undefined,
    }

    try {
        const response = await fetch(`${import.meta.env.VITE_API_DEVBOOK}${req_opt.path}`, options)

        if (response.ok) {
            return await response.json() as ResponseJsonSuccess
        } else {
            try {
                return await response.json() as ResponseJsonError
            } catch (error) {
                NotificationError('Error from frontend application, please try again later')
                return {
                    error: 'Error from frontend application',
                    errors: [error],
                } as ResponseJsonError
            }
        }
    } catch (error) {
        NotificationError('Error from frontend application or network error, please try again later')
        return {
            error: 'Error from frontend application or network error',
            errors: [error],
        } as ResponseJsonError
    }
}