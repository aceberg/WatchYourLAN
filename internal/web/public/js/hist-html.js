
function getHistHTML(hist) {

    hist = sortHistByDate(hist);

    let html = '', col, title;

    for (let h of hist) {
        if (h.Now != 0) {
            col = `fill:var(--bs-success);stroke:var(--bs-primary);`;
        } else {
            col = `fill:var(--bs-gray-500);stroke:var(--bs-primary);`;
        }
        title = `title="Date: ${h.Date}\nIface: ${h.Iface}\nIP: ${h.IP}\nKnown: ${h.Known}"`;

        html = html + `<i ${title}><svg width="10" height="20">
            <rect width="10" height="20" style="${col}"/>
            Sorry, your browser does not support inline SVG.  
        </svg></i>`;
    }
    return html;
}

function sortHistByDate(hist) {

    hist.sort((a, b) => (a.Date > b.Date));

    return hist;
}