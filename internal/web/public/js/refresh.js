let ref = false;
let refInterval;

autoRefresh();

function toggleRefresh() {
    ref = !ref;

    autoRefresh()
}

function autoRefresh() {

    if (ref) {
        const timeout = document.getElementById("ref-timeout").value;
        console.log("Refresh timeout", timeout);
        refInterval = setInterval(loadAddrs, timeout * 1000);
    } else {
        clearInterval(refInterval);
    }
}