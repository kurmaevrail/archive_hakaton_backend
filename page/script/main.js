document.addEventListener("DOMContentLoaded", main)

async function DBrow(target) {
    const response = await fetch(target);

    if (response.status === 200) {
        const data = await response.blob();
        const abuf = await data.arrayBuffer();
        const floats = new Float32Array(abuf)
        console.log("The data: " + floats)
    } else {
        console.log(`Error code ${response.status}`)
    }
}

function main() {
    console.log("Hello, world")
    const target = '/db/'
    DBrow(target + '0:1')
}
