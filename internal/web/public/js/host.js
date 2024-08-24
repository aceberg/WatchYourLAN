
async function delHost(id) {
    
    const url = '/api/host/del/'+id;

    await fetch(url);

    window.location.href = '/';
}

async function loadHistory(mac) {
    
    const url = '/api/history/'+mac;

    let hist = await (await fetch(url)).json();

    // console.log("HIST", hist);
    displayHistory(hist);
}

function displayHistory(hist) {

    let html, col, title;

    for (let h of hist) {
        if (h.Now != 0) {
            col = `fill:var(--bs-success);stroke:var(--bs-primary);`;
        } else {
            col = `fill:var(--bs-gray-500);stroke:var(--bs-primary);`;
        }
        title = `title="Date: ${h.Date}\nIface: ${h.Iface}\nIP: ${h.IP}\nKnown: ${h.Known}"`;

        html = `<i ${title}><svg width="10" height="20">
            <rect width="10" height="20" style="${col}"/>
            Sorry, your browser does not support inline SVG.  
        </svg></i>`;

        // html = `<i class="bi bi-file-fill" style="${col}" ${title}></i>`;
        document.getElementById('showHist').insertAdjacentHTML('beforeend', html);
    }
}