let oldFilter = '';
let bkpArray;

function filterFunc(field, value) {

    if (oldFilter == field) {
        addrsArray = bkpArray;
    }

    switch (field) {
        case 'iface':
            addrsArray = addrsArray.filter((item) => item.Iface == value);
            break;
        case 'known':
            addrsArray = addrsArray.filter((item) => item.Known == value);
            break;
        case 'line':
            addrsArray = addrsArray.filter((item) => item.Now == value);
            break;
        default:
            console.log("Filter error");
    }
    
    oldFilter = field;

    displayArrayData(addrsArray);
}

function resetFilter() {
    addrsArray = bkpArray;
    displayArrayData(addrsArray);
}

function searchFunc() {
    const s = document.getElementById('search').value;

    if (s != "") {

        const sl = s.toLowerCase();

        let newArray = [];

        for (let item of addrsArray) {
            
            if (searchItem(item, sl)) {
                newArray.push(item);
            }
        }
        addrsArray = newArray;
    } else {
        addrsArray = bkpArray;
    }

    displayArrayData(addrsArray);
}

function searchItem(item, sl) {

    const name = item.Name.toLowerCase();
    const hw = item.Hw.toLowerCase();
    const mac = item.Mac.toLowerCase();

    if ((name.includes(sl)) || (item.Iface.includes(sl)) || (item.IP.includes(sl)) || (mac.includes(sl)) || (hw.includes(sl)) || (item.Date.includes(sl))) {
        return true;
    } else {
        return false;
    }
}