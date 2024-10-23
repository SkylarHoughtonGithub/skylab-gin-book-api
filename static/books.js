const bookApiUrl = '/api/books';

async function listBooks() {
    try {
        const response = await fetch(bookApiUrl);
        
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        
        const data = await response.json();
        
        
        if (!data.books || !Array.isArray(data.books)) {
            throw new Error('Response does not contain an array of books');
        }
        
        const booksList = document.getElementById('books-list');
        booksList.innerHTML = '';
        
        data.books.forEach((book) => {
            const li = document.createElement('li');
            li.textContent = `${book.id}: ${book.title} by ${book.author}`;
            booksList.appendChild(li);
        });

        // Optionally, you can display the limit and offset
        console.log(`Showing ${data.books.length} books. Limit: ${data.limit}, Offset: ${data.offset}`);
    } catch (error) {
        console.error('Failed to list books:', error);
    }
}

async function addBook() {
    const title = document.getElementById('add-book-title').value;
    const author = document.getElementById('add-book-author').value;
    const newBook = { title, author };

    await fetch(bookApiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newBook),
    });

    ListBooks();
}

async function updateBook() {
    const index = document.getElementById('update-book-index').value;
    const title = document.getElementById('update-book-title').value;
    const author = document.getElementById('update-book-author').value;
    const updatedBook = { title, author };

    try {
        const response = await fetch(`${bookApiUrl}/${index}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedBook),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to update book');
        }

        const updatedBookData = await response.json();
        console.log('Book updated successfully:', updatedBookData);
        ListBooks();
    } catch (error) {
        console.error('Error updating book:', error.message);
    }
}

async function deleteBook() {
    const index = document.getElementById('delete-book-index').value;

    await fetch(`${bookApiUrl}/${index}`, {
        method: 'DELETE',
    });

    ListBooks();
}

// Load books when the page loads
window.onload = () => {
    listBooks();
};