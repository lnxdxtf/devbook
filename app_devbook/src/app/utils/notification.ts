import { useToast } from "vue-toastification";

const toast = useToast()

export const NotificationSuccess = (message: string) => {
    toast.success(message)
}

export const NotificationError = (message: string) => {
    toast.error(message)
}

export const NotificationInfo = (message: string) => {
    toast.info(message)
}

export const NotificationWarning = (message: string) => {
    toast.warning(message)
}

