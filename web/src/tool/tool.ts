import dayjs from "dayjs";

export function getBelongToText(val: string) {
    const number = parseInt(val);
    if (number === 0) {
        return '毫秒镜像'
    } else if (number === 1) {
        return 'Docker Hub'
    }
}


export function formatBytes(bytes: number) {
    if (bytes === 0) return '0 Bytes';

    const k = 1000;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}


export function getUpdateTime(last_updated: string) {

    if (last_updated == undefined ||last_updated == ''){
        return '';
    }
    // 获取当前时间
    const now = dayjs();
    // 指定的时间
    const specifiedTime = dayjs(last_updated);
    // 计算时间差
    const diffInMinutes = now.diff(specifiedTime,'minute');

    if (diffInMinutes < 1){
        return '刚刚更新'
    }

    const diffInHours = diffInMinutes / 60;

    if (diffInHours < 1){
        return diffInMinutes.toFixed(0) + '分钟前更新'
    }

    const diffInDays = diffInHours / 24;

    if (diffInDays < 1){
        return diffInHours.toFixed(0) + '小时前更新'
    }

    return diffInDays.toFixed(0) + '天前更新'

}