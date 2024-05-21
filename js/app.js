import { types } from "./cmds.js"
import { asciiPixel, createDecodeFrame, expand, pushFrame } from "./decode/frame.js"
import { parseFrame } from "./net/frame.js"
import { WS } from "./ws/index.js"

// TODO: provide a url?

/**
 * @param {HTMLElement} el
 */
function run(el) {
    const ws = new WS("ws://localhost:8080/ws")

    const decodeFrame = createDecodeFrame()
    ws.onMessage(async function(buf) {

        const frame = parseFrame(buf)
        switch (frame.cmd) {
        case types.open:
            console.log("OPEN", frame)
            break
        case types.frame:
            pushFrame(decodeFrame, frame)
            expand(decodeFrame)
            const pixels = asciiPixel(decodeFrame)

            // render
            // console.log(data.slice(0, 5))
            break
        default:
            throw new Error("unhandled frame type " + frame.cmd)
        }

    })
}

run(document.body)
