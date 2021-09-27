import random
import util
from datetime import date, timedelta
from itertools import repeat

titles = ['PSO Geo', 'PSO Region', 'PSO Territory', 'Engagement Name',
          'Engagement Id', 'Engagement Status', 'Engagement Type',
          'Business Unit', 'Bill To Customer', 'End Customer ',
          'Budgeted Hours Total', 'Planned Hours Total', 'Created On',
          'Payment Type', 'Rec Rev Type', 'Billing SKU',
          'Engagement Amount', 'Billing Currency', 'Total Billed-Redeemed',
          'Remaining Amount', 'Operation Analysts', 'Time Approver',
          'Billed-Redeemed by Date Range\n',
          'Remaining Forecast\nby Date Range', 'Forecast Hours this Qtr',
          'Forecast Revenue of Current Qtr',
          'Forecast Revenue  of Current  Qtr USD',
          'Total Billable hours as of  prior Qtr end',
          'Total Billable hours this Qtr', 'Mgmt Rev as of Prior Qtr',
          'Mgmt Rev as of Prior Qtr USD', 'Qtr Management Revenue To-Date',
          'Qtr Management Revenue To-date USD', 'Internal Hours - QTD',
          'Subbing Hours - QTD', 'Expenses Billed-Redeemed By Date Range',
          'Contract Status', 'PO Number', 'Billed-Redeemed by Date Range (USD)',
          'Last Time Entry Date', 'Status During Date Range',
          'Total Billed -Redeemed (USD)', 'Engagement Amount in USD',
          'Total Actual Hours', 'Hours Based Forecast Revenue of Current Qtr',
          'Hours Based Forecast Revenue of Current Qtr(USD)',
          'Planned Hrs for Qtr', 'Program ', 'Project Profile', 'SSM',
          'Project\nManager', 'Engagement\nClosed Date', 'Agreement Number',
          'Opportunity Number']

technology = ['VCF', 'vRA', 'vRO/vROPs', 'NSX', 'SRM', 'Horizon', 'WorkspaceONE',
             'Carbon Black']
services = ['Residency', 'Design and Deploy', 'Deploy', 'Services']

customers = ['Martian Defence Network', 'Alien Affairs and Barter',
             'Martian Romance', 'Uranus Trade and Investment',
             'Plutopac', 'Solar-System Federation Bank', 'Mercury Mining']

def rand_territory():
    states = []
    states.extend(repeat('NSW', 6))
    states.extend(repeat('ACT', 5))
    states.extend(repeat('VIC-TAS', 5))
    states.extend(repeat('SA', 2))
    states.append('WA')
    states.append('QLD-NT')
    return random.choice(states)

def rand_engagement_id(customer):
    return f'{customer[:4].upper()}{random.randint(1000, 9999)}'

def rand_status():
    status = []
    status.extend(repeat('Work in progress', 10))
    status.extend(repeat('Staged', 3))
    status.append('Inactive')
    return random.choice(status)

def rand_customer():
    return random.choice(customers)

def create_revenue_row():
    amount_multiplier = 100
    usd_conversion = 0.75
    customer = rand_customer()
    budgeted_hours = random.randint(18, 800)
    engagement_amount = budgeted_hours * amount_multiplier
    total_billed_redeemed = random.random() * engagement_amount
    remaining_amount = engagement_amount - total_billed_redeemed
    forecast = random.randint(0, budgeted_hours)

    engagement_name = util.rand_name(technology, services)
    row = []
    row.append('AP')                                                        #  0 PSO Geo
    row.append('ANZ')                                                       #  1 PSO Region
    row.append(rand_territory())                                            #  2 PSO Territory
    row.append(engagement_name)                                             #  3 Engagement Name
    row.append(rand_engagement_id(customer))                                #  4 Engagement Id
    row.append(rand_status())                                               #  5 Engagement Status
    row.append('Custom')                                                    #  6 Engagement Type
    row.append(random.choice(['Server', 'Desktop']))                        #  7 Business Unit
    row.append(customer)                                                    #  8 Bill To Customer
    row.append('')                                                          #  9 End Customer
    row.append(budgeted_hours)                                              # 10 Budgeted Hours Total
    row.append(budgeted_hours)                                              # 11 Planned Hours (same as budget)
    row.append(date.today() - timedelta(days=random.randint(7, 180)))       # 12 Created On
    row.append(random.choice(['Credit', 'Cash']))                           # 13 Payment Type
    row.append(random.choice(['Fixed Fee', 'T&M']))                         # 14 Rec Rev Type
    row.append('HLP-SV')                                                    # 15 Billing SKU
    row.append(engagement_amount)                                           # 16 Engagement Amount
    row.append('AUD')                                                       # 17 Billing Currency
    row.append(total_billed_redeemed)                                       # 18 Total Billed-Redeemed
    row.append(remaining_amount)                                            # 19 Remaining Amount
    row.append(util.rand_name(util.last_names, util.initials))              # 20 Operation Analysts
    row.append(util.rand_name(util.first_names, util.last_names))           # 21 Time Approver
    row.append(total_billed_redeemed * 0.5)                                 # 22 Billed-Redeemed by Date Range
    row.append(remaining_amount)                                            # 23 Remaining Forecast by Date Range
    row.append(forecast)                                                    # 24 Forecast Hours this Qtr
    row.append(forecast * amount_multiplier)                                # 25 Forecast Revenue of Current Qtr
    row.append(forecast * amount_multiplier * usd_conversion)               # 26 Forecast Revenue  of Current  Qtr USD
    row.append(random.randint(0, budgeted_hours // 2))                      # 27 Total Billable hours as of  prior Qtr end
    row.append(random.randint(0, budgeted_hours // 5))                      # 28 Total Billable hours this Qtr
    row.append(0)                                                           # 29 Mgmt Rev as of Prior Qtr
    row.append(0)                                                           # 30 Mgmt Rev as of Prior Qtr USD 
    row.append(0)                                                           # 31 Qtr Management Revenue To-Date
    row.append(0)                                                           # 32 Qtr Management Revenue To-date USD
    row.append(0)                                                           # 33 Internal Hours - QTD
    row.append(0)                                                           # 34 Subbing Hours - QTD
    row.append(0)                                                           # 35 Expenses Billed-Redeemed By Date Range
    row.append(random.choice(['SOW Countersigned', 'PCR Countersigned']))   # 36 Contract Status
    row.append(f'PO{random.randint(10000, 99999)}')                         # 37 PO Number
    row.append(0)                                                           # 38 Billed-Redeemed by Date Range (USD)
    row.append(date.today() - timedelta(days=random.randint(0, 7)))         # 39 Last Time Entry Date
    row.append('Active')                                                    # 40 Status During Date Range
    row.append(total_billed_redeemed * usd_conversion)                      # 41 Total Billed -Redeemed (USD)
    row.append(engagement_amount * usd_conversion)                          # 42 Engagement Amount in USD
    row.append(random.randint(0, budgeted_hours // 2))                      # 43 Total Actual Hours
    row.append(forecast * amount_multiplier)                                # 44 Hours Based Forecast Revenue of Current Qtr
    row.append(forecast * amount_multiplier * usd_conversion)               # 45 Hours Based Forecast Revenue of Current Qtr(USD)
    row.append(random.randint(0, budgeted_hours // 2))                      # 46 Planned Hrs for Qtr
    row.append('')                                                          # 47 Program
    row.append(random.choice(['', 'PR01', 'PR02']))                         # 48 Project Profile
    row.append(util.rand_name(util.first_names, util.last_names))           # 49 SSM
    row.append(util.rand_name(util.first_names, util.last_names))           # 50 Project Manager
    row.append('')                                                          # 51 Engagement Closed Date
    row.append(random.randint(1000000, 9999999))                            # 52 Agreement Number
    row.append(random.randint(1000000, 9999999))                            # 53 Opportunity Number
    return row

def mock_revenue(n):
    data = []
    for i in range(n):
        data.append(create_revenue_row())
    return data