let show = 500;
let addrsArray;

loadAddrs();

async function loadAddrs() {
    
    const n = localStorage.getItem("histShow");
    if (n != null) {
        show = n;
    }

    const url = '/api/all';
    addrsArray = await (await fetch(url)).json();

    loadHistory();
}

async function loadHistory() {
    
    let tr, td, url, hist;
    let i = 0;

    document.getElementById('showHist').innerHTML = '';

    for (let a of addrsArray) {
        url = '/api/history/'+a.Mac;
        hist = await (await fetch(url)).json();
        hist = sortHistByDate(hist);

        if (show > 0) {
            hist = hist.slice(0, show);
        }

        td = getHistHTML(hist); // hist-html.js
        i = i + 1;

        tr = `
        <tr>
            <td style="opacity: 45%;">${i}.</td>
            <td>
                <p>${a.Name}</p>
                <p><a href="http://${a.IP}" target="blank">${a.IP}</a></p>
                <p><a href="/host/${a.ID}">${a.Mac}</a></p>
            </td>
            <td>${td}</td>
        </tr>`;

        document.getElementById('showHist').insertAdjacentHTML('beforeend', tr);
    }
}

function showHist(n) {
    show = n;
    localStorage.setItem("histShow", show);
    loadHistory();
}

function manualShow() {
    const n = document.getElementById('man-show').value;
    showHist(n);
}