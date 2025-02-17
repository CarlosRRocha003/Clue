function fetchAllCards() {
    const url = 'http://localhost:8080/allCards';

    fetch(url)
        .then(response => response.json())
        .then(data => {
            populateGrid(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

let players = [];

function fetchAllPlayers() {
    const url = 'http://localhost:8080/allPlayers';
    players = [];
    fetch(url)
        .then(response => response.json())
        .then(data => {
            data.forEach((player) => players.push(player));
            populateComboBox2(players,"whoHasIt");
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

const Category = {
    Suspect: 0,
    Place: 1,
    Weapon: 2
};

function populateGrid(allCards) {
    const cardGrid = document.getElementById('cardGrid');
    
    cardGrid.innerHTML = '';
    allCards.sort((a, b) => a.Category - b.Category).forEach(card => {
        const categoryKey = Object.keys(Category).find(key => Category[key] === card.Category);
        const cardDiv = document.createElement('div');
        cardDiv.className = 'card';
        const whoHasItSelect = document.createElement('select');
        players.forEach(player => {
            const option = document.createElement('option');
            option.value = player.Id;
            option.text = player.Name;
            if (card.WhoHasIt.Id === player.Id) {
                option.selected = true;
            }
            whoHasItSelect.appendChild(option);
        });

        cardDiv.innerHTML = `
            <p>Name: ${card.Name}</p>
            <p>Category: ${categoryKey}</p>
            <p>Who Has It: </p>
        `;
        cardDiv.appendChild(whoHasItSelect);
        cardGrid.appendChild(cardDiv);
    });
}

function addPlayer() {
    const playerName = document.getElementById('playerNameInput').value;

    const url = 'http://localhost:8080/addPlayer';

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name: playerName })
    })
    .then(response => response.json())
    .then(data => {
        fetchAllPlayers();
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function addCard() {
    const cardName = document.getElementById('cardNameInput').value;
    const category = parseInt(document.getElementById('categorySelector').value);
    let card = { name: cardName, category: category}
    const url = 'http://localhost:8080/addCard';

    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(card)
    })
    .then(response => response.json())
    .then(data => {
        fetchAllCards();
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

function fetchCategoryList() {
    const url = 'http://localhost:8080/categories';

    fetch(url)
        .then(response => response.json())
        .then(data => {
            populateComboBox(data, 'categorySelector');
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function populateComboBox(data, comboboxId) {
    const selector = document.getElementById(comboboxId);

    selector.innerHTML = '';

    for (const key in data) {
        if (data.hasOwnProperty(key)) {
            const option = document.createElement('option');
            option.value = key;
            option.textContent = data[key];

            selector.appendChild(option);
        }
    }
}

function populateComboBox2(data, comboboxId) {
    const selector = document.getElementById(comboboxId);

    selector.innerHTML = '';

    data.forEach(obj => {
        const option = document.createElement('option');
        option.value = obj.Id;
        option.textContent = obj.Name;
        selector.appendChild(option);
    });
}

function objectToDictionary(obj) {
    let map = new Map();
    
    for (let key in obj) {
        if (obj.hasOwnProperty(key)) {
            map.set(key, obj[key]);
        }
    }
    
    return map;
}

window.onload = function() {
    fetchCategoryList();
    fetchAllPlayers();
};