/**
 * @param {String} header block header
 * @param {String} message block message
 * @param {int?} timeout when block will be removed
 * @return {HTMLDivElement}
 */
export function createSuccessBlock(header, message, timeout = 5000) {

    const id = "success-block-" + Math.random()

    let block = document.createElement('div')
    let blockHeader = document.createElement('div')
    let blockBody = document.createElement('div')

    blockHeader.innerText = header
    blockBody.innerText = message

    block.classList.add('success-message')
    blockHeader.classList.add('success-message-header')
    blockBody.classList.add('success-message-body')

    block.append(blockHeader)
    block.append(blockBody)
    block.setAttribute("id", id)

    setTimeout(() => {
        removeSuccessBlock(id)
    }, timeout)

    return block
}

/**
 * @param {String} id id of error block
 */
function removeSuccessBlock(id) {
    console.log("remove")
    document.getElementById(id).remove()
}