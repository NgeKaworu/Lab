module.exports = function (source) {
    console.log('source', source)

    // Apply some transformations to the source...
    const str = {
        source
    }
    return `export default ${JSON.stringify(str)}`
}
