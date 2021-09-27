import random
import util
from datetime import date

titles = ['PSO Geo','PSO Region','PSO Subregion','Engagement Id',
          'Engagement Name','Field Project Manager','Project Profile',
          'Budget Status (PMO)','Project Status (PMO)',
          'Customer Satisfaction (PMO)','Schedule Status (PMO)','Comments',
          'Mitigation Action','Engagement Manager','Practice Manager',
          'Budget Status UpdateOn','Budget Status Update By',
          'Customer Satisfaction UpdateOn','Customer Satisfaction Update By',
          'ScheduleStatus UpdateOn','Schedule Status Update By',
          'Project Status UpdateOn','Project Status Update By']

def create_gar_row(rev):
    row = []
    row.append('AP')                                                # PSO Geo
    row.append('ANZ')                                               # PSO Region
    row.append('Australia')                                         # PSO Subregion
    row.append(rev[4])                                              # Engagement Id
    row.append(rev[3])                                              # Engagement Name
    row.append(rev[50])                                             # Field Project Manager
    row.append(rev[48])                                             # Project Profile
    row.append(random.choice(['Green', 'Amber', 'Red']))            # Budget Status (PMO)
    row.append(random.choice(['Green', 'Amber', 'Red']))            # Project Status (PMO)
    row.append(random.choice(['Green', 'Amber', 'Red']))            # Customer Satisfaction (PMO)
    row.append(random.choice(['Green', 'Amber', 'Red']))            # Schedule Status (PMO)
    row.append('')                                                  # Comments
    row.append('')                                                  # Mitigation Action
    row.append(rev[21])                                             # Engagement Manager
    row.append(util.rand_name(util.first_names, util.last_names))   # Practice Manager
    row.append(date.today())                                        # Budget Status UpdateOn
    row.append(rev[50])                                             # Budget Status Update By
    row.append(date.today())                                        # Customer Satisfaction UpdateOn
    row.append(rev[50])                                             # Customer Satisfaction Update By
    row.append(date.today())                                        # ScheduleStatus UpdateOn
    row.append(rev[50])                                             # Schedule Status Update By
    row.append(date.today())                                        # Project Status UpdateOn
    row.append(rev[50])                                             # Project Status Update By
    return row

def mock_gar(revenue):
    data = []
    for rev in revenue:
        data.append(create_gar_row(rev))
    return data