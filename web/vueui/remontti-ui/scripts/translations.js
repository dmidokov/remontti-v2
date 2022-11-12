/**
 * @param {Array} translations The translations list
 * @param {String} label The label we are looking for
 */
export function getTranslations(translations, label) {
    if (translations === undefined) {
        return ""
    } else {
        return translations[label]
    }
}