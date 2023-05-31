const data = Deno.args[0]
const secret = Deno.args[1]
const signature = await sign(data, secret)
console.log("data: ", data)
console.log("secret: ", secret)
console.log("signature: ", signature)



async function sign(data: string, secret:string) {
    const encoder = new TextEncoder()
    const key = await crypto.subtle
        .importKey("raw", encoder.encode(secret), { name: "HMAC", hash: "SHA-256" }, false, ["sign"])
    const buff = await crypto.subtle.sign({ name: "HMAC", hash: "SHA-256"}, key, encoder.encode(data))
    return Array
        .from(new Uint8Array(buff))
        .map(i => i.toString(16).padStart(2, "0"))
        .join("")
}
