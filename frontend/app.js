const API_URL = 'http://localhost:8080';
let token = localStorage.getItem('token');

// Elementy DOM
const authModal = document.getElementById('authModal');
const loginForm = document.getElementById('loginForm');
const registerForm = document.getElementById('registerForm');
const loginBtn = document.getElementById('loginBtn');
const registerBtn = document.getElementById('registerBtn');
const logoutBtn = document.getElementById('logoutBtn');
const uploadSection = document.getElementById('uploadSection');
const uploadForm = document.getElementById('uploadForm');
const gallery = document.getElementById('gallery');

// Obsługa autoryzacji
function updateAuthUI() {
    if (token) {
        loginBtn.classList.add('hidden');
        registerBtn.classList.add('hidden');
        logoutBtn.classList.remove('hidden');
        uploadSection.classList.remove('hidden');
    } else {
        loginBtn.classList.remove('hidden');
        registerBtn.classList.remove('hidden');
        logoutBtn.classList.add('hidden');
        uploadSection.classList.add('hidden');
    }
}

function showToast(message, type = 'success') {
    const toast = document.getElementById('toast');
    toast.textContent = message;
    toast.className = `toast ${type}`;
    toast.classList.remove('hidden');
    setTimeout(() => toast.classList.add('hidden'), 3000);
}

// Obsługa modalu
loginBtn.onclick = () => {
    authModal.classList.remove('hidden');
    loginForm.classList.remove('hidden');
    registerForm.classList.add('hidden');
};

registerBtn.onclick = () => {
    authModal.classList.remove('hidden');
    loginForm.classList.add('hidden');
    registerForm.classList.remove('hidden');
};

document.querySelector('.close').onclick = () => {
    authModal.classList.add('hidden');
};

// Logowanie
async function login() {
    const nickname = document.getElementById('loginNickname').value;
    const password = document.getElementById('loginPassword').value;

    try {
        const response = await fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ nickname, password }),
        });

        const data = await response.json();

        if (response.ok) {
            token = data.token;
            localStorage.setItem('token', token);
            authModal.classList.add('hidden');
            updateAuthUI();
            showToast('Zalogowano pomyślnie');
            loadImages();
        } else {
            showToast(data.error, 'error');
        }
    } catch (error) {
        showToast('Błąd podczas logowania', 'error');
    }
}

// Rejestracja
async function register() {
    const nickname = document.getElementById('registerNickname').value;
    const password = document.getElementById('registerPassword').value;

    try {
        const response = await fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ nickname, password }),
        });

        const data = await response.json();

        if (response.ok) {
            showToast('Zarejestrowano pomyślnie');
            loginForm.classList.remove('hidden');
            registerForm.classList.add('hidden');
        } else {
            showToast(data.error, 'error');
        }
    } catch (error) {
        showToast('Błąd podczas rejestracji', 'error');
    }
}

// Wylogowanie
logoutBtn.onclick = () => {
    token = null;
    localStorage.removeItem('token');
    updateAuthUI();
    loadImages();
    showToast('Wylogowano pomyślnie');
};

// Przesyłanie zdjęć
uploadForm.onsubmit = async (e) => {
    e.preventDefault();

    const formData = new FormData();
    const fileInput = document.getElementById('imageFile');
    const isPublic = document.getElementById('isPublic').checked;

    formData.append('file', fileInput.files[0]);
    formData.append('public', isPublic);

    try {
        const response = await fetch(`${API_URL}/images/upload`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
            body: formData,
        });

        if (response.ok) {
            showToast('Zdjęcie przesłano pomyślnie');
            fileInput.value = '';
            loadImages();
        } else {
            const data = await response.json();
            showToast(data.error, 'error');
        }
    } catch (error) {
        showToast('Błąd podczas przesyłania zdjęcia', 'error');
    }
};

// Ładowanie zdjęć
async function loadImages() {
    try {
        const response = await fetch(`${API_URL}/images/public`);
        const images = await response.json();

        gallery.innerHTML = '';
        images.forEach(image => {
            const card = document.createElement('div');
            card.className = 'image-card';
            card.innerHTML = `
                <img src="${API_URL}/uploads/${image.name}" alt="${image.name}">
                <div class="image-info">
                    <p>${image.name}</p>
                </div>
            `;
            gallery.appendChild(card);
        });
    } catch (error) {
        showToast('Błąd podczas ładowania zdjęć', 'error');
    }
}

// Inicjalizacja
updateAuthUI();
loadImages();