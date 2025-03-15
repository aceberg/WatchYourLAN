let ref = false;
let refInterval;

autoRefresh();

function toggleRefresh() {
    ref = !ref;
    
    localStorage.setItem("refAuto", ref);

    autoRefresh();
}

function autoRefresh() {

    ref = JSON.parse(localStorage.getItem("refAuto"));
    document.getElementById("ref").checked = ref;

    if (ref) {
        const timeout = document.getElementById("ref-timeout").value;
        console.log("Refresh timeout", timeout);
        refInterval = setInterval(loadAddrs, timeout * 1000);
    } else {
        clearInterval(refInterval);
    }
}