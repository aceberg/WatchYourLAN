
function getHistHTML(hist) {

    let html = '', col, title;

    for (let h of hist) {
        if (h.Now != 0) {
            col = `my-box-on`;
        } else {
            col = `my-box-off`;
        }
        title = `title="Date: ${h.Date}\nIface: ${h.Iface}\nIP: ${h.IP}\nKnown: ${h.Known}"`;

        html = html + `<i ${title} class="${col}"></i>`;
    }
    return html;
}

function sortHistByDate(hist) {

    hist.sort((a, b) => (a.Date < b.Date ? 1 : -1));

    return hist;
}