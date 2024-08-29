let show = 500;
let mac;

async function delHost(id) {
    
    const url = '/api/host/del/'+id;

    await fetch(url);

    window.location.href = '/';
}

async function loadHistory(m) {

    const n = localStorage.getItem("hostShow");
    if (n != null) {
        show = n;
    }
    
    mac = m;
    const url = '/api/history/'+mac;

    let hist = await (await fetch(url)).json();
    hist = sortHistByDate(hist);
    if (show > 0) {
        hist = hist.slice(0, show);
    }

    // console.log("HIST", hist);
    displayHistory(hist);
}

function displayHistory(hist) {

    document.getElementById('showHist').innerHTML = getHistHTML(hist); // hist-html.js
}

function showHist(n) {
    show = n;
    localStorage.setItem("hostShow", show);
    loadHistory(mac);
}

function manualShow() {
    const n = document.getElementById('man-show').value;
    showHist(n);
}