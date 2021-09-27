import random
import util
from datetime import date, timedelta

titles = ['Expired Quarter', 'Expired Month', 'Credits Expired Date',
          'Credits Start Date', 'Geo', 'Region', 'Sub Region', 'Territory',
          'PSO Geo', 'Segment', 'PSO Region', 'PSO Territory', 'CSE', 'ESE',
          'Cumtomer Name', 'Customer Number SFDC', 'Sales Order ID',
          'MyLearn Account ID', 'Redemption LOB', 'Planned Redemption Date',
          'Redemption Status', 'Class Name', 'Project Manager',
          'Credit Status Summary', 'Engagement BU', 'Allocated Amount',
          'Allocated Number', 'Allocated Comments']


def rand_month():
    return random.choice(['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                          'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'])

def create_credit_row(rev):
    row= []
    row.append(random.choice(['22Q3', '22Q4', '23Q1', '23Q2']))         # Expired Quarter
    row.append(rand_month())                                            # Expired Month
    row.append(date.today() + timedelta(days=random.randint(30, 240)))  # Credits Expired Date
    row.append(date.today() - timedelta(days=random.randint(30, 240)))  # Credits Start Date
    row.append('APJ')                                                   # Geo
    row.append('ANZ')                                                   # Region
    row.append('ANZ-ENT-CEN')                                           # Sub Region
    row.append('ANZ-ENT-CEN-EAM3')                                      # Territory
    row.append('APJ')                                                   # PSO Geo
    row.append('ANZ')                                                   # Segment
    row.append('ANZ')                                                   # PSO Region
    row.append(rev[2])                                                  # PSO Territory
    row.append(util.rand_name(util.first_names, util.last_names))       # CSE
    row.append(util.rand_name(util.first_names, util.last_names))       # ESE
    row.append(rev[8])                                                  # Cumtomer Name
    row.append(f'CUST-{random.randint(1000000,9999999)}')               # Customer Number SFDC
    row.append(random.randint(1000000, 9999999))                        # Sales Order ID
    row.append(random.randint(1000000, 9999999))                        # MyLearn Account ID
    row.append('Consulting')                                            # Redemption LOB
    row.append(date.today())                                            # Planned Redemption Date
    row.append('HOLD')                                                  # Redemption Status
    row.append(rev[4])                                                  # Class Name
    row.append(rev[50])                                                 # Project Manager
    row.append('')                                                      # Credit Status Summary
    row.append(rev[7])                                                  # Engagement BU
    row.append(rev[16])                                                 # Allocated Amount
    row.append(0)                                                       # Allocated Number
    row.append('')                                                      # Allocated Comments
    return row

def mock_credits(revenue):
    data = []
    for rev in revenue:
        data.append(create_credit_row(rev))
    return data