import random

titles = ['ID', 'Engagement Id', 'Comment']

comments = ['Lorem ipsum dolor sit amet',
            'At vero eos et accusamus et iusto odio dignissimos',
            'Sed laoreet purus nibh, eget pretium sem viverra ac',
            'Donec venenatis porta quam id scelerisque',
            'Aenean cursus, purus eu porta tincidunt'
            'Ex diam mollis diam, in fermentum mauris odio vitae nisi',
            'In non lobortis lorem',
            'Integer imperdiet ante maximus est dignissim consequat',
            'Suspendisse potenti',
            'Etiam vel sem urna',
            'Etiam eu arcu id arcu porta varius ac et nibh',
            'Nulla diam lectus, fermentum eget luctus vel, elementum vel',
            'Quisque semper velit ac tristique egestas',
            'Suspendisse pulvinar ante non ligula convallis tincidunt in nec',
            'Fusce nec purus tellus']

def create_comment_row(id, engagemet_id):
    return [id, engagemet_id, random.choice(comments)]

def mock_comments(revenue, n):
    comments = []
    for i in range(n):
        idx = random.randint(0, len(revenue)-1)
        comments.append(create_comment_row(i+1, revenue[idx][4]))
    return comments
