export interface ResponseJsonSuccess {
    data: Object
}

export interface ResponseJsonError {
    error?: string | Object,
    errors?: string[] | Object | Object[]
}

export enum Method {
    GET = 'GET',
    POST = 'POST',
    PUT = 'PUT',
    DELETE = 'DELETE'
}


export interface RequestOptions {
    path: string,
    method: Method,
    headers?: HeadersInit,
    body?: BodyInit | Object
}

export async function fetchAPIDevBook(req_opt: RequestOptions): Promise<ResponseJsonSuccess | ResponseJsonError> {
    if (!req_opt.headers) {
        let baseHeaders: Headers
        baseHeaders = new Headers()
        baseHeaders.append('Content-Type', 'application/json')
        req_opt.headers = baseHeaders
    }

    const options = {
        method: req_opt.method,
        headers: req_opt.headers,
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
                return {
                    error: 'Error from frontend application',
                    errors: [error],
                } as ResponseJsonError
            }
        }
    } catch (error) {
        return {
            error: 'Error from frontend application or network error',
            errors: [error],
        } as ResponseJsonError
    }
}