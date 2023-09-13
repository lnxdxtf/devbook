import { Directive } from "vue";

export const timePassed: Directive<any> = {
    mounted: (el: any, binding: any) => {
        const datePassed = new Date(binding.value).getTime();
        const currentDate = new Date().getTime();

        const diffInMs = currentDate - datePassed;

        const diffInDays = Math.floor(diffInMs / (1000 * 60 * 60 * 24));
        const diffInWeeks = Math.floor(diffInDays / 7);
        const diffInMonths = Math.floor(diffInDays / 30.44);
        const diffInYears = Math.floor(diffInDays / 365.25);

        let formattedDiff;

        if (diffInYears >= 1) {
            formattedDiff = `${diffInYears} year${diffInYears > 1 ? 's' : ''} ago`;
        } else if (diffInMonths >= 1) {
            formattedDiff = `${diffInMonths} month${diffInMonths > 1 ? 'es' : ''} ago`;
        } else if (diffInWeeks >= 4) {
            formattedDiff = "1 month ago";
        } else if (diffInWeeks >= 1) {
            formattedDiff = `${diffInWeeks} week${diffInWeeks > 1 ? 's' : ''} ago`;
        } else {
            formattedDiff = `${diffInDays} day${diffInDays > 1 ? 's' : ''} ago`;
        }

        el.innerText = formattedDiff;
    }
}