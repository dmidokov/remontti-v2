/**
 * @param {String} header block header
 * @param {String} message block message
 * @param {int?} timeout when block will be removed
 * @return {HTMLDivElement}
 */
export function createErrorBlock(header, message, timeout = 5000) {

    const id = "error-block-" + Math.random()

    let block = document.createElement('div')
    let blockHeader = document.createElement('div')
    let blockBody = document.createElement('div')

    blockHeader.innerText = header
    blockBody.innerText = message

    block.classList.add('error-message')
    blockHeader.classList.add('error-message-header')
    blockBody.classList.add('error-message-body')

    block.append(blockHeader)
    block.append(blockBody)
    block.setAttribute("id", id)

    setTimeout(() => {
        removeErrorBlock(id)
    }, timeout)

    return block
}

/**
 * @param {String} id id of error block
 */
function removeErrorBlock(id) {
    console.log("remove")
    document.getElementById(id).remove()
}