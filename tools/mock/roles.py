import util

titles = ['Engagement Id','Project Manager','CSE','TAM','EAM']

def create_roles_row(rev):
    row = []
    row.append(rev[4])                                              # Engagement Id 
    row.append(rev[50])                                             # Project Manager
    row.append(rev[49])                                             # CSE
    row.append(util.rand_name(util.first_names, util.last_names))   # TAM
    row.append(util.rand_name(util.first_names, util.last_names))   # EAM
    return row

def mock_roles(revenue):
    data = []
    for rev in revenue:
        data.append(create_roles_row(rev))
    return data